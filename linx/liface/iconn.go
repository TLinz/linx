package liface

import "net"

type IConn interface {
	Start()
	Stop()
	GetConn() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
}
