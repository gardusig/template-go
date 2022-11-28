package main

import (
	"context"
	"flag"
	"log"

	"github.com/gardusig/template-go/example/grpc/client"
	"github.com/gardusig/template-go/example/grpc/model"
)

func main() {
	flag.Parse()
	client := client.NewGreeterClient("localhost:50051")
	r, err := client.Client.SayHello(
		context.Background(),
		&model.HelloRequest{
			Name: "world",
		},
	)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}
