package emailSender

import (
	rpcMessageCreator "emailSenderAPI/pkg/API/RPC/messageCreator"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"
)

// SendEmail sends the email with smtp method
func SendEmail(emails []string, message []byte) error {
	file, err := os.Open("pkg/smtp/emailSender/config.json")
	if err != nil {
		return fmt.Errorf("error opening config json file")
	}
	defer file.Close()
	var config rpcMessageCreator.Config
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
	timer := time.NewTimer(3 * time.Second)
	control := make(chan int)

	go smtpSend(control, smtpHost+":"+smtpPort, auth, email, emails, message, &err)
	select {
	case <-timer.C:
		log.Println("Timeout!")
		return fmt.Errorf("timeout")
	case <-control:
		break
	}
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error sending email")
	}
	return nil
}

func smtpSend(control chan int, addr string, auth smtp.Auth, from string, to []string, msg []byte, err *error) {
	*err = smtp.SendMail(addr, auth, from, to, msg)
	control <- 1
}
