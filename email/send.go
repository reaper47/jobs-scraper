package email

import (
	"bytes"
	"net/smtp"
	"text/template"

	"github.com/reaper47/jobs-scraper/config"
	"github.com/reaper47/jobs-scraper/model"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject string) *Request {
	return &Request{
		from:    config.GetEnvVar("from"),
		to:      to,
		subject: subject,
	}
}

func (r *Request) parseTemplate(fname string, data interface{}) error {
	t, err := template.ParseFiles(fname)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}

	r.body = buffer.String()
	return nil
}

func (r *Request) sendMail() error {
	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	auth := LoginAuth(config.GetEnvVar("username"), config.GetEnvVar("password"))
	if err := smtp.SendMail(config.GetEnvVar("smtp"), auth, r.from, r.to, []byte(body)); err != nil {
		return err
	}
	return nil
}

func (r *Request) Send(template string, jobs *model.Websites) error {
	if err := r.parseTemplate(template, jobs); err != nil {
		return err
	}

	if err := r.sendMail(); err != nil {
		return err
	}
	return nil
}
