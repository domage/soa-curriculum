package main

import (
	"context"
	"io"
	"log"

	"g_rpc/service"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:5050", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := service.NewGreeterClient(conn)

	// Contact the server and print out its response.
	// name := "WOlolo"
	// if len(os.Args) > 1 {
	// 	name = os.Args[1]
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := client.SayHello(ctx, &service.HelloRequest{Name: name})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }

	// log.Printf("Greeting: %s", resp.GetMessage())

	// countResp, err := client.Calculate(ctx, &service.CalculateRequest{A: 4, B: 5})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }

	// log.Printf("Result: %d", countResp.GetResult())

	// entities, err := client.CalculateHistory(ctx, &service.CalculateRequest{A: 1, B: 2})

	// bytes, _ := json.Marshal(entities)

	// log.Printf("Result:", string(bytes), err)

	stream, err := client.CalculateHistoryStream(context.Background(), &service.CalculateRequest{A: 1, B: 2})
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println(feature)
	}
}
