package REST

import (
	"emailSenderAPI/pkg/API/REST/messageCreator"
	"emailSenderAPI/pkg/DBlogging"
	"emailSenderAPI/pkg/smtp/emailSender"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)
// ClosurePost returns a function for our gin POST method
// hiding the database in it
func ClosurePost (dtb *sqlx.DB) func(*gin.Context) {
	db := dtb
	return func (c *gin.Context) {
	var msg restMessageCreator.JsonMessage
	msg.UnmarshalGin(c)
	to, message := msg.CreateMessage()
	err := emailSender.SendEmail(to, message)
	dbmsg, dbreq := msg.PrepToLog()
	if err != nil {
		dbreq.ResponseCode = "400"
		go DBlogging.LogRequest(db, dbmsg, dbreq)
		log.Println("Error sending message!\n", err)
		c.AbortWithError(400, err)
	}
	go DBlogging.LogRequest(db, dbmsg, dbreq)
	}
}
