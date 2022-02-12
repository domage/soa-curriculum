package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

var PORT = ":5454"

type Test struct {
	Qwe2 int64
	Qwe3 string
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
			fmt.Println("-> ", input.Qwe2, input.Qwe3)
		}

		time.Sleep(5 * time.Second)

		fmt.Println("Closing connection")
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
