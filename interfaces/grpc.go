package interfaces

import (
	"architecture/handlers/grpc"
	"architecture/message"
	"log"
	"net"

	gogprc "google.golang.org/grpc"
)

type GRpc struct {
	Service *message.Service
}

func (i GRpc) Run() {
	listen, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("Port 7777 is not available. %v", err)
	}

	gRpc := gogprc.NewServer()

	grpc.RegisterSmsServer(gRpc, grpc.NewSmsGRpc(i.Service))

	gRpc.Serve(listen)
}

func (i GRpc) IsBackground() bool {
	return true
}

func (i GRpc) Priority() int {
	return 257
}
