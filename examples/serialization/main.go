package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v2"
)

type Qwe struct {
	Rty string
}

type Test struct {
	ID         int
	Name       string
	ServiceIDs []int
	Tests      []Qwe
}

func writeToFile(name string, bytes []byte) {
	f, _ := os.Create(name)
	defer f.Close()

	f.Write(bytes)
	f.Sync()
}

func main() {
	t := Test{
		ID:         1,
		Name:       "Test",
		ServiceIDs: []int{1, 2, 3},
		Tests: []Qwe{
			Qwe{Rty: "rty"},
		},
	}

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	err := enc.Encode(t)

	netBytes := network.Bytes()

	var q Test
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, q.ID, q.ServiceIDs)

	jsonBytes, err := json.Marshal(t)

	fmt.Println("JSON:", string(jsonBytes))

	yamlBytes, err := yaml.Marshal(t)

	fmt.Println("YAML:", string(yamlBytes))

	msgBytes, err := msgpack.Marshal(t)

	fmt.Println("MsgPack:", string(msgBytes))

	writeToFile("./bytes.txt", netBytes)
	writeToFile("./json.txt", jsonBytes)
	writeToFile("./yaml.txt", yamlBytes)
	writeToFile("./msgpack.txt", msgBytes)

	fi, err := os.Open("./bytes.txt")
	defer fi.Close()

	var fromFile Test
	decoder := gob.NewDecoder(fi)
	err = decoder.Decode(&fromFile)

	fmt.Printf("%q: {%d, %d}\n", fromFile.Name, fromFile.ID, fromFile.ServiceIDs)
}
