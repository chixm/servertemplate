package server

import (
	"net/smtp"
)

// Mail sending function. only send, not to receive.
// Tested by Sending Gmail from Gmail server.

// tests by sending email
func InitializeEmailSender() {
	if Config.Email != nil && Config.Email.TestSendAddr != `` {
		Logger.Infoln(`Begin sending Email test`)
		to := []string{Config.Email.TestSendAddr}
		msg := []byte("Subject: test from ATAGO \r\n\r\n testing send mail")
		err := sendEmail(to, "template", msg)
		if err != nil {
			panic(err)
		}
		Logger.Infoln(`End sending Email test`)
	}
}

func sendEmail(to []string, from string, message []byte) error {
	auth := smtp.PlainAuth("", Config.Email.User, Config.Email.Password, Config.Email.Smtp)
	err := smtp.SendMail(Config.Email.SmtpSvr, auth, from, to, message)
	if err != nil {
		Logger.Errorln(err)
		return err
	}
	return nil
}
