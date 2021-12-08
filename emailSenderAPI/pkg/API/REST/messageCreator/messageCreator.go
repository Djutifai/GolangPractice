package messageCreator

func (m *JsonMessage) CreateMessage() ([]string, []byte) {
	var msg []byte
	msg = []byte("To: " + m.To + "\nFrom: " + m.To + "\nSubject: " +
		m.Subject + "\n\n" + m.Message)
	var to []string
	to = append(to, m.To)
	to = append(to, m.Copy...)
	return to, msg
}
