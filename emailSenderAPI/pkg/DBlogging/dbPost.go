package DBlogging

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"strings"
	"time"
)

type DBLog struct {
	Protocol     string
	MessageId    string
	ResponseCode string
}

// LogMessage inserting in our database the information about transaction.
// This method uses DBMessage and DBRequest as fillers for query
func LogMessage(db *sqlx.DB, msg *DBMessage, rq *DBLog) {
	if err := db.Ping(); err == nil {
		db.MustExec(schema)
		insertMsgQuery := `
		INSERT INTO message (from_email, to_email, subject, message, copy_to)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`
		insertRqQuery := `
		INSERT INTO request (protocol, message, response_code, created_at)
		VALUES ($1, $2, $3, $4)`
		msgId := 0
		err := db.Get(&msgId, insertMsgQuery, strings.ToLower(msg.From), msg.To, msg.Subject,
			msg.Message, pq.Array(msg.Cc))
		if err != nil {
			log.Println("Error logging request!\n", err)
		}
		if err != nil {
			log.Println("Error getting last inserted id in messages!\n", err)
		}
		_, err = db.Exec(insertRqQuery, rq.Protocol, msgId, rq.ResponseCode,
			time.Now().Format(`2006-01-02 15:04:05`))
		if err != nil {
			log.Println("Error logging request!\n", err)
		}
	} else {
		log.Println("Still no connection to DB!")
	}
}
