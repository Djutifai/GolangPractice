package DBlogging

import (
	"github.com/jmoiron/sqlx"
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
);

CREATE INDEX IF NOT EXISTS from_email_index ON message (from_email);
`

type DBMessage struct {
	From    string
	To      string
	Subject string
	Message string
	Cc      []string
}

// StartDb connecting to psql db and creates if not exists the needed table
func StartDb() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=server password=server dbname=logger sslmode=disable")
	for err != nil {
		log.Println(err)
		time.Sleep(3 * time.Second)
		db, err = sqlx.Connect("postgres", "host=postgres port=5432 user=server password=server dbname=logger sslmode=disable")
	}
	db.MustExec(schema)
	return db
}
