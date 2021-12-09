package emailSender

import (
	"emailSenderAPI/pkg/API/REST/messageCreator"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendEmail(emails []string, message []byte) error {
	file, err := os.Open("../pkg/smtp/emailSender/config.json")
	if err != nil {
		return fmt.Errorf("error opening config json file")
	}
	defer file.Close()

	var config messageCreator.Config
	err = config.Unmarshal(file)
	if err != nil {
		log.Println(err)
		return err
	}
	// Sender email and password
	email := config.SenderConfig["email"]
	pass := config.SenderConfig["pass"]
	// Destination emails
	// smtp server configuration
	smtpHost := config.SmtpConfig["host"]
	smtpPort := config.SmtpConfig["port"]
	// Making authenticate variable
	auth := smtp.PlainAuth("", email, pass, smtpHost)
	// sending our message to emails
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, email, emails, message)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error sending email")
	}
	return nil
}
