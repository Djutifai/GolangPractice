package gRpc

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	UnimplementedMsgHandlerServer
}

func (s *server) SendMessage(ctx context.Context, in *Message) (*ErrHandler, error) {
	log.Printf("Made a deal with the devil\n")
	return &ErrHandler{
		Err: "All good",
	}, nil
}

func MakeServer () {
	port := flag.Int("port", 8081, "The server port")
	flag.Parse()
	net.Listen("TCP", fmt.Sprintf("localhost:%d", port))
	g := grpc.NewServer()
	RegisterMsgHandlerServer(g, &server{})
	fmt.Println("Server is done")
}
