package gRpc

/* Researching atm how to transfer select from psql in a fancy way via grpc
import (
	"emailSenderAPI/pkg/DBlogging"
	"fmt"
	"golang.org/x/net/context"
	"strings"
)

func (g *GRPCServer) Get(ctx context.Context, msg *GETMessage) (*GETResponse, error) {
	if msg.FromMail == "" {
		return &GETResponse{RespCode: "400", RespMessage: "No id provided by user!"}, nil
	}
	res, err := DBlogging.LogRequestFromMail(g.Db, msg.FromMail)
	if err != nil {
		return &GETResponse{RespCode: "500", RespMessage: "Some error occured on our server!"}, err
	} else if len(res) == 0 {
		return &GETResponse{RespCode: "204", RespMessage: "No messages was found from " + msg.FromMail}, nil
	}
	respMsg := dbGetReturnTo(res)
	return &GETResponse{RespCode: "200", RespMessage: respMsg}, nil
}

func dbGetReturnTo(res []DBlogging.DBGetReturn) string {
	var str string
	for _, elem := range res {
		str = fmt.Sprintf("{\nFrom: %s\nTo: %s\nSubject: %s\n"+
			"Message: %s\nCopy to %s\nProtocol: %s\nResponseCode: %s\nCreated at: %s\n}\n",
			elem.From, elem.To, elem.Subject, elem.Message, strings.Join(elem.Cc, ", "), elem.Protocol,
			elem.ResponseCode, elem.CreatedAt)
	}
	return str
}
*/
