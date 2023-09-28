package mail

import "github.com/DeniesKresna/gobridge/serror"

type IMail interface {
	SystemSendMail(receiver string, subject string, content string) serror.SError
}
