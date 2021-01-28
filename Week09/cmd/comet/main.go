package main

import (
	"fmt"
	"github.com/neverhover/Go-001/tree/main/Week09/pkg/comet"
	"net"
	"os"
)

var (
	bindAddr = ":10000"
)

func main() {
	listen, err := net.Listen("tcp", bindAddr)
	if err != nil {
		fmt.Printf("Can't start server on %s. error is %s", bindAddr, err)
		os.Exit(1)
	}
	fmt.Printf("Start tcp server on %s\n", bindAddr)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Can't accept connection. error is %s", err)
			continue
		}
		go comet.Handle(conn)
	}
}
