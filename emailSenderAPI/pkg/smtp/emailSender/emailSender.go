package emailSender

import (
	"emailSenderAPI/pkg/API/REST/messageCreator"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// SendEmail sends the email with smtp method
func SendEmail(emails []string, message []byte) error {
	file, err := os.Open("../pkg/smtp/emailSender/config.json")
	if err != nil {
		return fmt.Errorf("error opening config json file")
	}
	defer file.Close()

	var config restMessageCreator.Config
	err = config.UnmarshalConfig(file)
	if err != nil {
		log.Println(err)
		return err
	}
	email := config.SenderConfig["email"]
	pass := config.SenderConfig["pass"]
	smtpHost := config.SmtpConfig["host"]
	smtpPort := config.SmtpConfig["port"]
	auth := smtp.PlainAuth("", email, pass, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, email, emails, message)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error sending email")
	}
	return nil
}
