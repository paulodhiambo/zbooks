package main

import (
	"google.golang.org/grpc"
	"userservice/internal"
	"userservice/pkg"
)

func main() {
	//grpc Server
	grpcServer := grpc.NewServer()

	// Initialize the application
	application := internal.Initialize(grpcServer)

	// Start the web server
	pkg.StartServer(application, grpcServer)
}
