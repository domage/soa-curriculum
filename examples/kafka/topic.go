package main

import (
	"net"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	// to create topics when auto.create.topics.enable='false'
	topic := "test2"

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	var controllerConn *kafka.Conn
	controllerConn, err := dialer.Dial("tcp", net.JoinHostPort("localhost", "29092"))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     2,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}
