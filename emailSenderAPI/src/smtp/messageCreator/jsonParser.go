package messageCreator

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

type Config struct {
	SenderConfig map[string]string `json:"senderConfig"`
	SmtpConfig map[string]string `json:"smtpConfig"`
}
type Message struct {
	From    string   `json:"from"`
	To      string   `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Copy    []string `json:"copy"`
}

func (c *Config) Unmarshal (file *os.File) error {
	byte, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Error in reading file\n")
	}
	err = json.Unmarshal(byte, c)
	if err != nil {
		return fmt.Errorf("Error in unmarshalling config json file\n")
	}
	return nil
}

func (m *Message) Unmarshal (c *gin.Context) error {
	body := c.Request.Body
	msg, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("Error reading body message from request\n")
	}
	err = json.Unmarshal(msg, m)
	if err != nil {
		return fmt.Errorf("Error in unmarshalling body message\n")
	}
	return nil
}

