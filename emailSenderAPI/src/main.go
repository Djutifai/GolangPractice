package main

import (
	"emailSenderAPI/pkg/API/REST"
	"emailSenderAPI/pkg/API/gRpc"
	"fmt"
	"github.com/gin-gonic/gin"
	grpc2 "google.golang.org/grpc"
	"log"
	"net"
	//"fmt"
	//"github.com/gin-gonic/gin"
	//"net/http"
)

// SMTP - done!
// right now im semi done with grpc
func main() {
	r := gin.Default()
	r.POST("/sendmsg", REST.PostMsg)
	go r.Run("localhost:8080")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8081"))
	if err != nil {
		log.Fatal(err)
	} else {
		grpcServer := grpc2.NewServer()
		srv := &gRpc.GRPCServer{}
		gRpc.RegisterSendMessageServer(grpcServer, srv)
		fmt.Println("Running grpcServer")
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Grpc done")
	}
}
