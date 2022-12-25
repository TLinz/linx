package linet

import (
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

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}

			// Basic display business.
			go func() {
				for {
					buf := make([]byte, 512)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Read err:", err)
						continue
					}

					fmt.Printf("Recv client: %s\n", buf)

					_, err = conn.Write(buf[:n])
					if err != nil {
						fmt.Println("Write err:", err)
						continue
					}
				}
			}()
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
