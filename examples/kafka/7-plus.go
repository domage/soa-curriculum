package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var topic = "result"

type Action struct {
	Action int `json:"action"`
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "plus",
		"auto.offset.reset": "earliest",
	})

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092"})
	if err != nil {
		panic(err)
	}

	defer c.Close()
	defer p.Close()

	c.SubscribeTopics([]string{"plus"}, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Request for plus:", msg)
		fmt.Println("Request for plus:", string(msg.Value))

		act := &Action{}
		json.Unmarshal(msg.Value, act)

		bytes, _ := json.Marshal(&Action{
			Action: 0,
			Value1: act.Value1 + act.Value2,
			Value2: act.Value2,
		})

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          bytes,
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
