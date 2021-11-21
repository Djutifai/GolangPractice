package main

import (
	"emailSenderAPI/src/smtp/emailSender"
	"github.com/gin-gonic/gin"

	//"fmt"
	//"github.com/gin-gonic/gin"
	//"net/http"
)
// SMTP - done!
// right now im semi done with grpc
func main () {
	r := gin.Default()
	r.POST("/sendmsg", emailSender.PostMsg)
	//gRpc.MakeServer()
	r.Run("localhost:8080")
}
