package mailhandle

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"

	"github.com/go-mail/mail"
)

type MailHandler struct {
	From      string
	FromName  string
	EmailHost string
	EmailPort int
}

type MailInfo struct {
	From       string   `josn:"from,omitempty"`
	FromName   string   `josn:"from_name,omitempty"`
	To         string   `josn:"to"`
	Subject    string   `josn:"subject"`
	Body       string   `json:"body"`
	Attachment []string `josn:"attachments,omitempty"`
}

func NewMailHandler() *MailHandler {
	return &MailHandler{
		From:      "test@example.com",
		FromName:  "Tester",
		EmailHost: "localhost",
		EmailPort: 1025,
	}
}

func (h *MailHandler) SendMail(info MailInfo) error {
	m := mail.NewMessage()
	if info.From == "" || info.From == "null" {
		m.SetHeader("From", h.From)
	} else {
		m.SetHeader("From", info.From)
	}

	if info.FromName == "" || info.FromName == "null" {
		m.SetHeader("FromName", h.FromName)
	} else {
		m.SetHeader("FromName", info.FromName)
	}

	m.SetHeader("To", info.To)
	m.SetHeader("Subject", info.Subject)

	parsedBody, err := h.ParseTemplate("./cmd/templates/mail.html.gohtml", info)
	if err != nil {
		return err
	}
	fmt.Println(parsedBody)
	m.SetBody("text/html", parsedBody)

	if len(info.Attachment) > 0 {
		for _, attach := range info.Attachment {
			m.Attach(attach)
		}
	}

	d := mail.NewDialer(h.EmailHost, h.EmailPort, "jim", "")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = d.DialAndSend(m)

	if err != nil {
		return err
	}

	return nil
}

func (h *MailHandler) ParseTemplate(templateFileName string, info MailInfo) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	buf := new(bytes.Buffer)
	err = t.ExecuteTemplate(buf, "Body", info)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Printf("parsed string is: %s", buf.String())
	return buf.String(), nil
}
