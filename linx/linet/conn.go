package linet

import (
	"fmt"
	"linx/liface"
	"net"
)

type handler func(*net.TCPConn, []byte, int) error

type Conn struct {
	conn     *net.TCPConn
	connID   uint32
	isClosed bool
	handler  handler
}

func NewConn(conn *net.TCPConn, connID uint32, handler handler) liface.IConn {
	return &Conn{
		conn:     conn,
		connID:   connID,
		isClosed: false,
		handler:  handler,
	}
}

func (c *Conn) handleRead() {
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		n, err := c.conn.Read(buf)
		if err != nil {
			fmt.Println("Read err:", err)
			continue
		}

		fmt.Printf("Recv Conn[%d]: %s\n", c.connID, buf)

		if err = c.handler(c.conn, buf, n); err != nil {
			fmt.Printf("Conn[%d] handler err: %s\n", c.connID, err)
			break
		}
	}
}

func (c *Conn) Start() {
	// Deal with read business.
	go c.handleRead()
	// TODO: deal with write business.
}

func (c *Conn) Stop() {
	fmt.Printf("Stop Conn[%d]...\n", c.connID)
	if c.isClosed {
		return
	}
	c.isClosed = true
	if err := c.conn.Close(); err != nil {
		fmt.Printf("Close Conn[%d] err: %s\n", c.connID, err)
		return
	}
}

func (c *Conn) GetConn() *net.TCPConn {
	return c.conn
}

func (c *Conn) GetConnID() uint32 {
	return c.connID
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
