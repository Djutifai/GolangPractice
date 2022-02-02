package DBlogging

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"strings"
)

type DBGetReturn struct {
	From         string
	To           string
	Subject      string
	Message      string
	Cc           []string
	Protocol     string
	ResponseCode string
	CreatedAt    string
}

func LogRequestFromMail(db *sqlx.DB, fromMail string) ([]DBGetReturn, error) {
	selectQuery := `
	SELECT m.from_email, m.to_email, m.subject, m.message, m.copy_to,
	       r.protocol, r.response_code, r.created_at
	FROM message as m
	JOIN request as r
	ON r.message=m.id
	WHERE from_email=$1
	`
	rows, err := db.Query(selectQuery, strings.ToLower(fromMail))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	returnArray := make([]DBGetReturn, 0)
	for rows.Next() {
		getReturn := DBGetReturn{}
		err := rows.Scan(&getReturn.From, &getReturn.To, &getReturn.Subject,
			&getReturn.Message, pq.Array(&getReturn.Cc), &getReturn.Protocol,
			&getReturn.ResponseCode, &getReturn.CreatedAt)
		if err != nil {
			log.Println(err)
			return returnArray, err
		}
		returnArray = append(returnArray, getReturn)
	}
	return returnArray, nil
}
