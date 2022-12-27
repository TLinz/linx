package linet

import (
	"errors"
	"fmt"
	"net"

	"linx/liface"
)

type Server struct {
	Name string
	IPV  string
	IP   string
	Port int
}

func callBack(conn *net.TCPConn, buf []byte, n int) error {
	if _, err := conn.Write(buf[:n]); err != nil {
		fmt.Println("CallBack: Write err:", err)
		return errors.New("CallBack error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("Server [%s] starting...\n", s.Name)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPV, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Resolve TCP addr err:", err)
			return
		}

		listener, err := net.ListenTCP(s.IPV, addr)
		if err != nil {
			fmt.Println("Listen TCP err:", err)
			return
		}

		connID := 0

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}

			c := NewConn(conn, uint32(connID), callBack)
			go c.Start()
			connID++
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	// TODO: additional business

	select {}
}

func NewServer(name string) liface.IServer {
	return &Server{
		Name: name,
		IPV:  "tcp4",
		IP:   "0.0.0.0",
		Port: 8080,
	}
}
