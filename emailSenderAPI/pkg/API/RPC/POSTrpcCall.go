package RPC

import (
	"emailSenderAPI/pkg/API/RPC/messageCreator"
	"emailSenderAPI/pkg/DBlogging"
	"emailSenderAPI/pkg/smtp/emailSender"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

// ClosurePost returns a function for our gin POST method
// hiding the database in it
func ClosurePost(dtb *sqlx.DB) func(*gin.Context) {
	db := dtb
	return func(ctx *gin.Context) {
		var msg rpcMessageCreator.JsonMessage
		msg.UnmarshalGin(ctx)
		to, message := msg.CreateMessage()
		err := emailSender.SendEmail(to, message)
		dbmsg, dblog := msg.PrepToLog()
		if err != nil {
			checkError(ctx, db, err, dbmsg, dblog)
			return
		}
		if err := db.Ping(); err == nil {
			go DBlogging.LogMessage(db, dbmsg, dblog)
		}
	}
}

func checkError(ctx *gin.Context, db *sqlx.DB, err error, dbmsg *DBlogging.DBMessage, dbreq *DBlogging.DBLog) {
	checkErr := fmt.Sprintf("%v", err)
	if checkErr == "timeout" {
		dbreq.ResponseCode = "408"
		if err := db.Ping(); err == nil {
			go DBlogging.LogMessage(db, dbmsg, dbreq)
		}
		log.Println("Timeout!\n", err)
		ctx.AbortWithError(408, err)
	} else {
		dbreq.ResponseCode = "400"
		if err := db.Ping(); err == nil {
			go DBlogging.LogMessage(db, dbmsg, dbreq)
		}
		log.Println("Error sending message!\n", err)
		ctx.AbortWithError(400, err)
	}
}
