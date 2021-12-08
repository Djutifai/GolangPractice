package gRpc

import (
	"context"
	"emailSenderAPI/pkg/smtp/emailSender"
	"log"
)

type GRPCServer struct{}

func (g *GRPCServer) mustEmbedUnimplementedSendMessageServer() {
	panic("implement me")
}

func (g *GRPCServer) Send(ctx context.Context, msg *GMessage) (*Response, error) {
	to, message := msg.CreateMessage()
	err := emailSender.SendEmail(to, message)
	var res *Response
	if err != nil {
		log.Fatal(err)
		res.ErrCode = "500"
		res.RespMsg = "There was an error sending your mail!\n"
	} else {
		res.ErrCode = "200"
		res.RespMsg = "Message sent successfully!\n"
	}
	return res, nil
}

func (x *GMessage) CreateMessage() ([]string, []byte) {
	var msg []byte
	msg = []byte("To: " + x.To + "\nFrom: " + x.To + "\nSubject: " +
		x.Subject + "\n\n" + x.Msg)
	var to []string
	to = append(to, x.To)
	to = append(to, x.Cc...)
	return to, msg
}
