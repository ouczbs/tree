package netutil

import (
	"net"
)
type Flushable interface {
	Flush() error
}

type Connection struct {
	net.Conn
	//Flushable
}

func (n Connection) Flush() error {
	return nil
}

func NewConnection(conn net.Conn) * Connection{
	return &Connection{
		conn,
	}
}