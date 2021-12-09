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
	if err != nil {
		log.Println(err)
		return &Response{ErrCode: "500", RespMsg: "Error in sending email"}, err
	} else {
		return &Response{ErrCode: "200", RespMsg: "Email has been successfully sent!"}, nil
	}
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
