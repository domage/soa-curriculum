package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("publisher")
	opts.OnConnect = connectHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Publish("test/some-random-room/temperature", 0, false, "24.5")
	client.Publish("test/room/time", 0, false, "01:00")
	client.Publish("test/room2/time", 0, false, "01:00")
}
