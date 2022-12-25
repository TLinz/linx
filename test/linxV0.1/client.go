package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp4", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Connect server err:", err)
		return
	}

	for {
		_, err := conn.Write([]byte("Hello from the other side..."))
		if err != nil {
			fmt.Println("Write err:", err)
			return
		}

		buf := make([]byte, 512)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}

		fmt.Printf("Server back: %s\n", buf)

		time.Sleep(1 * time.Second)
	}
}
