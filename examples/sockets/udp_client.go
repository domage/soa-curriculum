package main

import (
	"fmt"
	"net"
)

func main() {
	CONNECT := ":5454"

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	_, err = c.Write([]byte("data"))

	buffer := make([]byte, 1)
	n, from, err := c.ReadFromUDP(buffer) // !
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Bytes transmitted:", n)
	fmt.Printf("Reply: %s\n", string(buffer))
	fmt.Printf("...was read from: %s\n", from)
}
