package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/hamba/avro"
	"github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v2"

	"serialization/serialization/models"
)

type Qwe struct {
	Rty string
}

type Test struct {
	ID         int    `json:"id" avro:"ID"`
	Name       string `json:"name" avro:"Name"`
	ServiceIDs []int  `json:"service_ids" yaml:"sids" avro:"ServiceIDs"`
	Tests      []Qwe  `json:"tests"`
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

	fmt.Println("-----")
	fmt.Println("JSON:", string(jsonBytes))

	yamlBytes, err := yaml.Marshal(t)

	fmt.Println("-----")
	fmt.Println("YAML:", string(yamlBytes))

	msgBytes, err := msgpack.Marshal(t)
	// gzip

	fmt.Println("-----")
	fmt.Println("MsgPack:", string(msgBytes))

	writeToFile("./bytes.txt", netBytes)
	writeToFile("./json.txt", jsonBytes)
	writeToFile("./yaml.txt", yamlBytes)
	writeToFile("./msgpack.txt", msgBytes)

	// AVRO
	schemaStr, err := ioutil.ReadFile("schema.avsc")
	if err != nil {
		panic(err)
	}

	schema, err := avro.Parse(string(schemaStr))
	data, err := avro.Marshal(schema, t)

	if err != nil {
		panic(err)
	}
	fmt.Println("Avro: ", string(data))

	out, err := proto.Marshal(&models.Person{
		Id:   int32(t.ID),
		Name: t.Name,
	})
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	fmt.Println("Protobuf: ", string(out))

	writeToFile("./avro.txt", data)
	writeToFile("./pb.txt", out)

	fi, err := os.Open("./bytes.txt")
	defer fi.Close()

	var fromFile Test
	decoder := gob.NewDecoder(fi)
	err = decoder.Decode(&fromFile)

	fmt.Printf("%q: {%d, %d}\n", fromFile.Name, fromFile.ID, fromFile.ServiceIDs)
}
