package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var totalTopic = "total"

type Action struct {
	Action int `json:"action"`
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "mult",
		"auto.offset.reset": "earliest",
	})

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092"})
	if err != nil {
		panic(err)
	}

	defer c.Close()
	defer p.Close()

	c.SubscribeTopics([]string{"result"}, nil)

	for {
		msg, _ := c.ReadMessage(-1)

		fmt.Println("Request for multi:", string(msg.Value))

		act := &Action{}
		json.Unmarshal(msg.Value, act)

		bytes, _ := json.Marshal(&Action{
			Action: 0,
			Value1: act.Value1 * act.Value2,
			Value2: act.Value2,
		})

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &totalTopic, Partition: kafka.PartitionAny},
			Value:          bytes,
		}, nil)
	}

	p.Flush(15 * 1000)
}
