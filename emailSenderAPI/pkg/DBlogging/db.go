package DBlogging

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var schema = `

CREATE TABLE IF NOT EXISTS message (
    id serial primary key,
    from_email varchar(255),
    to_email varchar(255),
    subject varchar(255),
    message text,
    copy_to varchar(255)[]
);

CREATE TABLE IF NOT EXISTS request (
		id serial primary key,
		protocol varchar(20),
		message integer references message(id),
		response_code smallint,
		created_at timestamp
);`

type DBlog struct {
	Id				int 	`db:"id"`
	Protocol		string 	`db:"protocol"`
	Message 		string	`db:"message"`
	ResponseCode	int		`db:"response_code"`
	CreatedAt 		string	`db:"timestamp"`
}

type DBRequest struct {
	Protocol 		string
	MessageId 		string
	ResponseCode	string
}

type DBMessage struct {
	From 	string
	To 		string
	Subject string
	Message string
	Cc 		[]string
}
// StartDb connecting to psql db and creates if not exists the needed table
func  StartDb() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=server dbname=logger sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.MustExec(schema)
	return db
}

// LogRequest inserting in our database the information about transaction.
// This method uses DBMessage and DBRequest as fillers for query
func LogRequest (db *sqlx.DB, msg *DBMessage, rq *DBRequest) {
	insertMsgQuery := `
	INSERT INTO message (from_email, to_email, subject, message, copy_to)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	insertRqQuery := `
	INSERT INTO request (protocol, message, response_code, created_at)
	VALUES ($1, $2, $3, $4)`
	msgId := 0
	err := db.Get(&msgId, insertMsgQuery, msg.From, msg.To, msg.Subject,
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
}