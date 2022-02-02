package gRpc

import (
	"emailSenderAPI/pkg/DBlogging"
	"emailSenderAPI/pkg/smtp/emailSender"
	"golang.org/x/net/context"
	"log"
)

// Send creates smtp format message from GMessage, calls emailSender.SendEmail
// and log to database with GMessage.PrepToLog.
func (g *GRPCServer) Post(ctx context.Context, msg *POSTMessage) (*POSTResponse, error) {
	to, message := msg.CreateMessage()
	err := emailSender.SendEmail(to, message)
	dbmsg, dbreq := msg.PrepToLog()
	if err != nil {
		log.Println("Error sending email with grpc!\n", err)
		dbreq.ResponseCode = "400"
		if err := g.Db.Ping(); err == nil {
			DBlogging.LogMessage(g.Db, dbmsg, dbreq)
		}
		return &POSTResponse{RespCode: "400", RespMsg: "Error in sending email"}, err
	}
	if err := g.Db.Ping(); err == nil {
		go DBlogging.LogMessage(g.Db, dbmsg, dbreq)
	}
	return &POSTResponse{RespCode: "200", RespMsg: "Email has been successfully sent!"}, nil
}
