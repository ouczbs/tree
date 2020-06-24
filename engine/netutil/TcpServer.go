package netutil

import (
	"github.com/ouczbs/tree/engine/gwlog"
	"github.com/ouczbs/tree/engine/ioutil"
	"net"
	"time"
)

const (
	_RESTART_TCP_SERVER_INTERVAL = 3 * time.Second
	_RESTART_UDP_SERVER_INTERVAL = 3 * time.Second
)

// TCPServerDelegate is the implementations that a TCP server should provide
type TCPServerDelegate interface {
	ServeTCPConnection(net.Conn)
}

// ServeTCPForever serves on specified address as TCP server, for ever ...
func ServeTCPForever(listenAddr string, delegate TCPServerDelegate) {
	for {
		err := serveTCPForeverOnce(listenAddr, delegate)
		gwlog.Errorf("server@%s failed with error: %v, will restart after %s", listenAddr, err, _RESTART_TCP_SERVER_INTERVAL)
		time.Sleep(_RESTART_TCP_SERVER_INTERVAL)
	}
}

func serveTCPForeverOnce(listenAddr string, delegate TCPServerDelegate) error {
	defer func() {
		if err := recover(); err != nil {
			gwlog.TraceError("serveTCPImpl: paniced with error %s", err)
		}
	}()

	return ServeTCP(listenAddr, delegate)

}

// ServeTCP serves on specified address as TCP server
func ServeTCP(listenAddr string, delegate TCPServerDelegate) error {
	ln, err := net.Listen("tcp", listenAddr)
	gwlog.Infof("Listening on TCP: %s ...", listenAddr)

	if err != nil {
		return err
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			if ioutil.IsTimeoutError(err) {
				continue
			} else {
				return err
			}
		}

		gwlog.Infof("Connection from: %s", conn.RemoteAddr())
		go delegate.ServeTCPConnection(conn)
	}
}
