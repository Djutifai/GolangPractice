package gRpc

import "emailSenderAPI/pkg/DBlogging"

// CreateMessage creates smtp message from GMessage struct
func (m *GMessage) CreateMessage() ([]string, []byte) {
	var msg []byte
	msg = []byte("To: " + m.To + "\nFrom: " + m.To + "\nSubject: " +
		m.Subject + "\n\n" + m.Message)
	var to []string
	to = append(to, m.To)
	to = append(to, m.Cc...)
	return to, msg
}
// PrepToLog creating and filling up structs that will be used for logging in database.
func (m *GMessage) PrepToLog () (*DBlogging.DBMessage, *DBlogging.DBRequest) {
	dbmsg := DBlogging.DBMessage{
		From:    m.From,
		To:      m.To,
		Subject: m.Subject,
		Message: m.Message,
		Cc:      m.Cc,
	}
	dbreq := DBlogging.DBRequest{
		Protocol:      "GRPC",
		ResponseCode:  "200",
	}
	return &dbmsg, &dbreq
}