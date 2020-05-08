package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"

	"personal/micro-demo/interface/proto/example"
)

type Example struct{}

func (e *Example) Hello(ctx context.Context, req *example.Request, rsp *example.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	example.RegisterExampleHandler(service.Server(), new(Example))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
