package REST

import (
	"emailSenderAPI/pkg/API/REST/messageCreator"
	"emailSenderAPI/pkg/smtp/emailSender"
	"fmt"
	"github.com/gin-gonic/gin"
)

func PostMsg(c *gin.Context) {
	var msg messageCreator.JsonMessage
	msg.UnmarshalGin(c)
	to, message := msg.CreateMessage()
	err := emailSender.SendEmail(to, message)
	if err != nil {
		c.AbortWithError(400, err)
		fmt.Printf("%v", err)
	}
}
