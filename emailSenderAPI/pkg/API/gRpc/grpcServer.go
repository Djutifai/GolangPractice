package gRpc

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	grpc2 "google.golang.org/grpc"
	"log"
	"net"
)

// GRPCServer is a struct needed for our proto-generated files.
// Contains with sqlx.DB
type GRPCServer struct {
	Db *sqlx.DB
}

func (g *GRPCServer) mustEmbedUnimplementedSendMessageServer() {
	panic("implement me")
}

// StartGrpcServer creates local GRPC server on 8081 port
func StartGrpcServer(db *sqlx.DB) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":8081"))
	defer lis.Close()
	if err != nil {
		log.Println(err)
	} else {
		grpcServer := grpc2.NewServer()
		srv := &GRPCServer{Db: db}
		RegisterSendMessageServer(grpcServer, srv)
		defer grpcServer.Stop()
		fmt.Println("Running grpcServer")
		if err = grpcServer.Serve(lis); err != nil {
			log.Println(err)
		}
		fmt.Println("Grpc done")
	}
}
