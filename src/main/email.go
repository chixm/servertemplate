package main

import (
	"net/smtp"
)

// Mail sending function. only send, not to receive.
// Tested by Sending Gmail from Gmail server.

// tests by sending email
func initializeEmailSender() {
	if config.Email != nil && config.Email.TestSendAddr != `` {
		logger.Infoln(`Begin sending Email test`)
		to := []string{config.Email.TestSendAddr}
		msg := []byte("Subject: test from ATAGO \r\n\r\n testing send mail")
		err := sendEmail(to, "template", msg)
		if err != nil {
			panic(err)
		}
		logger.Infoln(`End sending Email test`)
	}
}

func sendEmail(to []string, from string, message []byte) error {
	auth := smtp.PlainAuth("", config.Email.User, config.Email.Password, config.Email.Smtp)
	err := smtp.SendMail(config.Email.SmtpSvr, auth, from, to, message)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}
