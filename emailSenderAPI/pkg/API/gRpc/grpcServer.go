package gRpc

import (
	"context"
	"emailSenderAPI/pkg/DBlogging"
	"emailSenderAPI/pkg/smtp/emailSender"
	"github.com/jmoiron/sqlx"
	"log"
)
// GRPCServer is a struct needed for our proto-generated files.
// Contains with sqlx.DB
type GRPCServer struct{
	Db *sqlx.DB
}

func (g *GRPCServer) mustEmbedUnimplementedSendMessageServer() {
	panic("implement me")
}
// Send creates smtp format message from GMessage, calls emailSender.SendEmail
// and log to database with GMessage.PrepToLog.
func (g *GRPCServer) Send(ctx context.Context, msg *GMessage) (*Response, error) {
	to, message := msg.CreateMessage()
	err := emailSender.SendEmail(to, message)
	dbmsg, dbreq := msg.PrepToLog()
	if err != nil {
		log.Println("Error sending email with grpc!\n",err)
		dbreq.ResponseCode = "400"
		go DBlogging.LogRequest(g.Db, dbmsg, dbreq)
		return &Response{RespCode: "400", RespMsg: "Error in sending email"}, err
	}
	go DBlogging.LogRequest(g.Db, dbmsg, dbreq)
	return &Response{RespCode: "200", RespMsg: "Email has been successfully sent!"}, nil
}