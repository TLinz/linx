package main

import "linx/linet"

func main() {
	s := linet.NewServer("linxV0.1")
	s.Serve()
}
