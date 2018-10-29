package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("%s", err)
	}

	x := listener.Addr().String()
	words := strings.Split(x, ":")

	fmt.Println(words[1])
}
