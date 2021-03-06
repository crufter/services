package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/services/location/handler"
	"github.com/micro/services/location/ingester"
	proto "github.com/micro/services/location/proto/location"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.location"),
	)

	service.Init()

	proto.RegisterLocationHandler(service.Server(), new(handler.Location))

	micro.RegisterSubscriber(ingester.Topic, service.Server(), new(ingester.Geo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
