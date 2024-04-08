package main

import (
	"braces.dev/errtrace"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

// Initialize the MDNS service
func initMDNS(peerhost host.Host, rendezvous string) (<-chan peer.AddrInfo, error) {
	n := &discoveryNotifee{
		PeerChan: make(chan peer.AddrInfo),
	}
	svc := mdns.NewMdnsService(peerhost, rendezvous, n)
	if err := svc.Start(); err != nil {
		return nil, errtrace.Wrap(err)
	}
	return n.PeerChan, nil
}

type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

// interface to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}
