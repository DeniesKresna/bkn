package mail

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/DeniesKresna/gohelper/utstring"
	gomail "gopkg.in/mail.v2"
)

type Gmail struct {
	SystemMail     string
	SystemPassword string
	SMTPServer     string
	SMTPPort       int
}

func InitGmail() (IMail, error) {
	portStr := utstring.GetEnv(models.EmailPortEnv, "587")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return &Gmail{
		SystemMail:     utstring.GetEnv(models.EmailAddressEnv, "batkorumbawamedsos@gmail.com"),
		SystemPassword: utstring.GetEnv(models.EmailPasswordEnv, "ofubvzpswfrwbttb"),
		SMTPServer:     utstring.GetEnv(models.EmailServerEnv, "smtp.gmail.com"),
		SMTPPort:       port,
	}, nil
}

func (g *Gmail) SystemSendMail(receiver string, subject string, content string) (errx serror.SError) {
	functionName := "[Gmail.SystemSendMail]"

	m := gomail.NewMessage()

	m.SetHeader("From", g.SystemMail)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(g.SMTPServer, g.SMTPPort, g.SystemMail, g.SystemPassword)

	if err := d.DialAndSend(m); err != nil {
		utlog.Error(err)
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While DialAndSend Email", functionName))
		return
	}

	return
}
