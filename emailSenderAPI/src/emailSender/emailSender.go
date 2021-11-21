package emailSender

import (
	"emailSenderAPI/src/smtp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/smtp"
	"os"
)

func PostMsg(c *gin.Context) {
	var msg smtp.Message
	msg.Unmarshal(c)
	to, message := msg.CreateMessage()
	err := sendEmail(to, message)
	if err != nil {
		c.AbortWithError(400, err)
		fmt.Printf("%v", err)
	}
}

func sendEmail(emails []string, message []byte) error {
	file, err := os.Open("emailSender/config.json")
	if err != nil {
		return fmt.Errorf("Error opening config json file\n")
	}
	defer file.Close()

	var config smtp.Config
	err = config.Unmarshal(file)

	// Sender email and password
	email := config.SenderConfig["email"]
	pass :=  config.SenderConfig["pass"]
	// Destination emails
	// smtp server configuration
	smtpHost := config.SmtpConfig["host"]
	smtpPort := config.SmtpConfig["port"]
	// Making authenticate variable
	auth := smtp.PlainAuth("", email, pass, smtpHost)
	// sending our message to emails
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, email, emails, message)
	if err != nil {
		return fmt.Errorf("Error sending email!\n")
	}
	return nil
}