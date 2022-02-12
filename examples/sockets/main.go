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
	Qwe string
}

func main() {
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	// c, err := l.Accept()

	// fmt.Println("Client connected")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // !!!1

	// netData, err := bufio.NewReader(c).ReadString('\n')

	// if netData != "" {
	// 	fmt.Print("-> ", string(netData))
	// 	c.Write([]byte("Message received"))
	// }

	// !!!2

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		input := Test{}

		netData, err := bufio.NewReader(c).ReadBytes('\n')

		err = json.Unmarshal(netData, &input)

		fmt.Println("Raw string ", string(netData))
		fmt.Println(err)
		fmt.Println("-> ", input.Qwe)
		c.Write([]byte("Message received"))

		time.Sleep(5 * time.Second)

		fmt.Println("Closing connection")
	}
}
