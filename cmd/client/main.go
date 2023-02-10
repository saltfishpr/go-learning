package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const count = 500

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
	conns := make([]*Client, count)
	for i := 0; i < count; i++ {
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Printf("dial %d: %v", i, err)
			continue
		}
		log.Printf("connected, client id: %d", i)
		conns[i] = &Client{id: i, conn: c}
	}

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	for _, c := range conns {
		if c == nil {
			continue
		}
		wg.Add(1)
		go c.Start(ctx, &wg)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-sig:
		cancel()
	case <-wait(&wg):
		log.Println("all clients disconnected")
	}
}

func wait(wg *sync.WaitGroup) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)
		wg.Wait()
	}()

	return done
}

type Client struct {
	id   int
	conn *websocket.Conn
}

func (c *Client) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		case <-done:
			return
		case t := <-ticker.C:
			err := c.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("client %d: %s", c.id, t.String())))
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}
