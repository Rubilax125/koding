// Package sender provides an API for mail sending operations
package emailsender

// Mail struct hold the required parameters for sending an email
type Mail struct {
	To      string
	ToName  string
	Subject string
	Text    string
	// TODO maybe we can remove this HTML field
	HTML       string
	From       string
	Bcc        string
	FromName   string
	ReplyTo    string
	Properties *Properties
}

type Properties struct {
	Username string
	Options  map[string]interface{}
}

func NewProperties() *Properties {
	options := map[string]interface{}{}
	return &Properties{Username: "", Options: options}
}

func NewMail(to, from, subject, username string) *Mail {
	properties := NewProperties()
	properties.Username = username

	return &Mail{To: to, From: from, Subject: subject, Properties: properties}
}

func NewEmptyMail() *Mail {
	return &Mail{}
}

func (m Mail) GetId() int64 {
	return 0
}

func (m Mail) BongoName() string {
	return "api.mail"
}
