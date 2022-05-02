package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"g_rpc/models"
	"g_rpc/service"

	"google.golang.org/grpc"
)

type server struct {
	service.UnimplementedGreeterServer
}

func (server) SayHello(context.Context, *service.HelloRequest) (*service.HelloReply, error) {
	res := service.HelloReply{
		Message: fmt.Sprintf("Hello there, someone!"),
	}

	return &res, nil
}

func (server) Calculate(ctx context.Context, req *service.CalculateRequest) (*service.CalculateReply, error) {
	res := &service.CalculateReply{
		Result: req.A + req.B,
	}

	model := models.Calc{
		A:   req.A,
		B:   req.B,
		Res: res.Result,
	}

	_, err := models.PG.NewInsert().Model(&model).Exec(context.Background())

	return res, err
}

func (server) CalculateHistory(ctx context.Context, in *service.CalculateRequest) (*service.CalculateEntities, error) {
	modelList := []*models.Calc{}

	err := models.PG.NewSelect().Model(&modelList).Scan(context.Background())

	output := []*service.CalculateEntity{}

	for _, entity := range modelList {
		output = append(output, &service.CalculateEntity{
			Id:  entity.ID,
			A:   entity.A,
			B:   entity.B,
			Res: entity.Res,
		})
	}

	return &service.CalculateEntities{
		Entities: output,
	}, err
}

func (server) CalculateHistoryStream(req *service.CalculateRequest, stream service.Greeter_CalculateHistoryStreamServer) error {
	calc1 := &models.Calc{
		ID:  1,
		A:   1,
		B:   2,
		Res: 3,
	}
	calc2 := &models.Calc{
		ID:  1,
		A:   1,
		B:   2,
		Res: 3,
	}
	modelList := []*models.Calc{
		calc1, calc2,
	}

	// models.PG.NewSelect().Model(&modelList).Scan(context.Background())

	go func() {
		for i, entity := range modelList {
			fmt.Println("Sent", i)
			stream.Send(&service.CalculateEntity{
				Id:  entity.ID,
				A:   entity.A,
				B:   entity.B,
				Res: entity.Res,
			})
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Func ended")

	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	models.Init()

	fmt.Println("Socket bound")

	s := grpc.NewServer()
	service.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Service is alive")
}
