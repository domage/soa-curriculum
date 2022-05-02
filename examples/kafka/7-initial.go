package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var initialTopic = "plus"

type Action struct {
	Action int `json:"action"`
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "initial",
		"auto.offset.reset": "earliest",
	})

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	bytes, _ := json.Marshal(&Action{
		Action: 0,
		Value1: 7,
		Value2: 2,
	})

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &initialTopic, Partition: kafka.PartitionAny},
		Value:          bytes,
	}, nil)

	defer c.Close()
	c.SubscribeTopics([]string{"total"}, nil)

	for {
		msg, _ := c.ReadMessage(-1)
		act := &Action{}
		json.Unmarshal(msg.Value, act)

		fmt.Println("Total: ", act.Value1)
	}
}
