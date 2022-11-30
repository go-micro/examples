package main

import (
	"github.com/go-micro/examples/auth/handler"
	pb "github.com/go-micro/examples/auth/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	// Import the JWT auth plugin
	_ "github.com/go-micro/plugins/v4/auth/jwt"
)

var (
	name    = "helloworld"
	version = "latest"
)

func main() {

	// Create service
	srv := micro.NewService()

	srv.Init(
		micro.Name(name),
		micro.Version(version),
		micro.WrapHandler(NewAuthWrapper(srv)),
	)

	// Register handler
	if err := pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

}
