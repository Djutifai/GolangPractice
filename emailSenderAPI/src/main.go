package main

import (
	"emailSenderAPI/pkg/API/REST"
	"emailSenderAPI/pkg/API/gRpc"
	"emailSenderAPI/pkg/DBlogging"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	grpc2 "google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db := DBlogging.StartDb()
	go StartRestServer(db)
	StartGrpcServer(db)
}
// StartRestServer creates local REST server on 8080 port
func StartRestServer(db *sqlx.DB) {
	r := gin.Default()
	r.POST("/sendmsg", REST.ClosurePost(db))
	err := r.Run("localhost:8080")
	if err != nil {
		log.Println("Error creating REST server!\n", err)
	}
}
// StartGrpcServer creates local GRPC server on 8081 port
func StartGrpcServer(db *sqlx.DB) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8081"))
	if err != nil {
		log.Println(err)
	} else {
		grpcServer := grpc2.NewServer()
		srv := &gRpc.GRPCServer{Db: db}
		gRpc.RegisterSendMessageServer(grpcServer, srv)
		fmt.Println("Running grpcServer")
		if err = grpcServer.Serve(lis); err != nil {
			log.Println(err)
		}
		fmt.Println("Grpc done")
	}
}