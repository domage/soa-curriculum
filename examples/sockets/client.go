package main

import (
	"fmt"
	"net"
)

var SERVER_ADDRESS = "127.0.0.1:5454"

// func handler(c net.Conn) {
// 	netData, _ := bufio.NewReader(c).ReadBytes('\n')

// 	fmt.Println(string(netData))
// }

func main() {
	c, err := net.Dial("tcp", SERVER_ADDRESS)
	if err != nil {
		fmt.Println(err)
		return
	}

	// handler(c)
	// { sensors: [{number: "AA:BB:CC", value: 25.6}] }

	fmt.Fprintf(c, "{\"Qwe\": \"123ggg\"}\n")

	c.Close()
	// fmt.Fprintf(c, "sdfjhdsklfjhs")
}
