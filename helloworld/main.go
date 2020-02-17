package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/services/helloworld/handler"

	helloworld "github.com/micro/services/helloworld/proto/helloworld"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.helloworld"),
		micro.Version("latest"),
	)

	service.Init()

	// Register Handler
	helloworld.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
