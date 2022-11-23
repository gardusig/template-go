package main

import (
	"context"
	"flag"
	"log"

	"github.com/gardusig/template-go/pkg/greeter_service/client"
	"github.com/gardusig/template-go/pkg/greeter_service/model"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	client := client.NewGreeterClient(*addr)
	r, err := client.Client.SayHello(context.Background(), &model.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}
