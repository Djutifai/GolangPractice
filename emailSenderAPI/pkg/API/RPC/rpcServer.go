package RPC

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

// StartRestServer creates local RPC server on 8080 port
func StartRestServer(db *sqlx.DB) {
	r := gin.Default()
	r.POST("/sendMsg", ClosurePost(db))
	r.GET("/getMsg/:fromMail", ClosureGet(db))
	err := r.Run(":8080")
	if err != nil {
		log.Println("Error creating RPC server!\n", err)
	}
}
