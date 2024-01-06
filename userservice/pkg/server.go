package pkg

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"userservice/internal"
)

func StartServer(app *internal.Application, grpcServer *grpc.Server) {
	// listener address
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		app.Logger.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}
	// start serving to the address
	log.Fatal(grpcServer.Serve(lis))
}
