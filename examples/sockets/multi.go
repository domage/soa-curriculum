package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

var PORT = ":5454"

type Test struct {
	Qwe string
}

func handler(c net.Conn) {
	fmt.Println("Handler started")

	defer func() { fmt.Println("Handler ended") }()

	defer func() {
		fmt.Println("Client disconnected")
	}()
	c.Write([]byte("Client connected\n"))

	for {
		input := Test{}

		netData, _ := bufio.NewReader(c).ReadBytes('\n')
		err := json.Unmarshal(netData, &input)

		if err == nil {
			fmt.Print("-> ", input.Qwe)
		}
	}
}

func main() {
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handler(c)
	}
}
