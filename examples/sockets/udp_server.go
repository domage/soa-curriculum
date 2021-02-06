package main

import (
	"fmt"
	"net"
)

func main() {
	PORT := ":5454"

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))
		fmt.Printf("...send from %s\n", addr)

		data := []byte("wololo")
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr) // !
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
