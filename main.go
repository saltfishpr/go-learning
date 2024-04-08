package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"fmt"
	"log/slog"
	"maps"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"braces.dev/errtrace"
	"github.com/labstack/echo/v4"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
	"github.com/sourcegraph/conc"
)

func main() {
	cfg, err := initConfig()
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	if err != nil {
		fatal(logger, err)
	}

	sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d", cfg.ListenHost, cfg.ListenPort))
	if err != nil {
		fatal(logger, err)
	}
	logger.Info(fmt.Sprintf("Listening on: %s", sourceMultiAddr.String()))
	host, err := libp2p.New(
		libp2p.Identity(prvKey),
		libp2p.ListenAddrs(sourceMultiAddr),
	)
	if err != nil {
		fatal(logger, err)
	}

	app, err := NewApp(cfg, logger, host)
	if err != nil {
		fatal(logger, err)
	}
	go app.Start(ctx)
	app.StartBackend(ctx)
}

func fatal(logger *slog.Logger, err error) {
	logger.Error(fmt.Sprintf("%+v", err))
	os.Exit(1)
}

type Conn struct {
	logger *slog.Logger
	stream network.Stream
	done   chan struct{}
}

func NewConn(baseLogger *slog.Logger, stream network.Stream) *Conn {
	return &Conn{
		logger: baseLogger.With(slogAttrPeerID(stream.Conn().RemotePeer())),
		stream: stream,
		done:   make(chan struct{}),
	}
}

func (c Conn) RemoteID() peer.ID {
	return c.stream.Conn().RemotePeer()
}

func (c *Conn) Close() error {
	if c.IsClosed() {
		return errtrace.Errorf("conn already closed")
	}

	close(c.done)
	return c.stream.Reset()
}

func (c *Conn) Send(message string) error {
	if c.IsClosed() {
		return errtrace.Errorf("conn already closed")
	}

	w := bufio.NewWriter(bufio.NewWriter(c.stream))
	if _, err := w.WriteString(message); err != nil {
		return errtrace.Wrap(err)
	}
	if err := w.Flush(); err != nil {
		return errtrace.Wrap(err)
	}
	return nil
}

func (c *Conn) Recv() {
	ch := c.recv()
	for {
		select {
		case <-c.done:
			return
		case message, ok := <-ch:
			if !ok {
				return
			}
			c.logger.Info(message)
		}
	}
}

func (c *Conn) recv() <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		r := bufio.NewReader(bufio.NewReader(c.stream))
		for {
			message, err := r.ReadString('\n')
			if err != nil {
				c.logger.Error("failed to read from peer", slog.Any("err", err))
				return
			}
			ch <- message
		}
	}()
	return ch
}

func (c *Conn) IsClosed() bool {
	select {
	case <-c.done:
		return true
	default:
	}
	return false
}

type App struct {
	cfg       *config
	logger    *slog.Logger
	host      host.Host // libp2p host, 既是服务端也是客户端
	protocol  protocol.ID
	peerCh    <-chan peer.AddrInfo
	messageCh chan string

	connWg conc.WaitGroup
	connMu sync.Mutex
	conns  map[peer.ID]*Conn
}

func NewApp(cfg *config, logger *slog.Logger, host host.Host) (*App, error) {
	peerCh, err := initMDNS(host, cfg.Rendezvous)
	if err != nil {
		return nil, err
	}

	a := &App{
		cfg:       cfg,
		logger:    logger,
		host:      host,
		protocol:  protocol.ID(cfg.Protocol),
		peerCh:    peerCh,
		messageCh: make(chan string),
		conns:     make(map[peer.ID]*Conn),
	}

	return a, nil
}

func (a *App) StartBackend(ctx context.Context) {
	e := echo.New()

	e.POST("/chat", func(c echo.Context) error {
		type Req struct {
			Message string `json:"message"`
		}

		reqBody := new(Req)
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		a.messageCh <- reqBody.Message
		return c.NoContent(http.StatusOK)
	})

	e.GET("/statistics", func(c echo.Context) error {
		conns := a.cloneConnMap()
		return c.JSON(http.StatusOK, echo.Map{
			"peersCount": len(conns),
		})
	})

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", a.cfg.BackendPort)); err != nil && err != http.ErrServerClosed {
			fatal(a.logger, err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		fatal(a.logger, err)
	}
}

func (a *App) Start(ctx context.Context) {
	a.host.SetStreamHandler(a.protocol, a.handleStream)

	for {
		select {
		case <-ctx.Done():
			a.shutdown()
			return
		case peer := <-a.peerCh:
			if peer.ID > a.host.ID() {
				continue
			}
			if err := a.connectPeer(ctx, peer); err != nil {
				a.logger.Error("failed to connect to peer", slogAttrPeerID(peer.ID), slog.Any("err", err))
			}
		case message := <-a.messageCh:
			a.sendMessage(ctx, message)
		}
	}
}

func (a *App) shutdown() {
	for _, conn := range a.cloneConnMap() {
		if err := conn.Close(); err != nil {
			a.logger.Error("failed to close conn", slogAttrPeerID(conn.RemoteID()), slog.Any("err", err))
		}
		a.delConn(conn)
	}
	a.connWg.Wait()
}

// handleStream 处理对端连接.
func (a *App) handleStream(stream network.Stream) {
	a.logger.Info("connecting from peer", slogAttrPeerID(stream.Conn().RemotePeer()))

	if _, ok := a.conns[stream.Conn().RemotePeer()]; ok {
		a.logger.Info("peer already connected", slogAttrPeerID(stream.Conn().RemotePeer()))
		if err := stream.Reset(); err != nil {
			a.logger.Error("failed to reset stream", slogAttrPeerID(stream.Conn().RemotePeer()), slog.Any("err", err))
		}
		return
	}

	conn := NewConn(a.logger, stream)
	go a.startConn(conn)
}

// connectPeer 发现 peer 主动连接.
func (a *App) connectPeer(ctx context.Context, peer peer.AddrInfo) error {
	// check peer is already connected

	a.logger.Info("connecting to peer", slogAttrPeerID(peer.ID))
	if err := a.host.Connect(ctx, peer); err != nil {
		return errtrace.Wrap(err)
	}

	stream, err := a.host.NewStream(ctx, peer.ID, a.protocol)
	if err != nil {
		return errtrace.Wrap(err)
	}

	conn := NewConn(a.logger, stream)
	go a.startConn(conn)

	return nil
}

func (a *App) startConn(conn *Conn) {
	a.connWg.Go(func() {
		a.addConn(conn)
		defer a.delConn(conn)

		conn.Recv()
	})
}

func (a *App) sendMessage(ctx context.Context, message string) {
	for _, conn := range a.cloneConnMap() {
		select {
		case <-ctx.Done():
			return
		default:
		}
		a.logger.Error("send message to peer", slogAttrPeerID(conn.RemoteID()), slog.String("message", message))
		if err := conn.Send(message); err != nil {
			a.logger.Error("failed to send message to peer", slogAttrPeerID(conn.RemoteID()), slog.Any("err", err))
		}
	}
}

func (a *App) addConn(conn *Conn) {
	a.connMu.Lock()
	defer a.connMu.Unlock()
	a.conns[conn.RemoteID()] = conn
}

func (a *App) delConn(conn *Conn) {
	a.connMu.Lock()
	defer a.connMu.Unlock()
	delete(a.conns, conn.RemoteID())
}

func (a *App) cloneConnMap() map[peer.ID]*Conn {
	a.connMu.Lock()
	defer a.connMu.Unlock()
	return maps.Clone(a.conns)
}

func slogAttrPeerID(id peer.ID) slog.Attr {
	return slog.String("peer", id.String())
}
