package main

import (
	"fmt"
	"net/smtp"
)

const (
	user = "AKIAIKPNOS3WJXVHATCQ"
	pwd  = "An1t90naxXpgSaoZeQBiHEqlLDzPv1C1ZbVy3XEIlyhs"

	name = "開發團隊"
	from = "service@daily-beauty.xyz"
	to   = "pudding850806@gmail.com"

	host = "email-smtp.us-west-2.amazonaws.com"
	port = 587
)

func connect() smtp.Auth {
	a := smtp.PlainAuth("", user, pwd, host)
	return a
}

func createMsg(html string) []byte {
	msg := ""
	msg += fmt.Sprintf("To: %s\r\n", to)
	msg += fmt.Sprintf("From: %s <%s>\r\n", name, from)
	msg += fmt.Sprintf("Subject: %s\r\n", "Subject HERE")

	msg += "MIME-version: 1.0;\r\n"
	msg += `Content-Type: text/html; charset="UTF-8"` + "\r\n"

	msg += "\r\n"
	msg += fmt.Sprintf("%s\r\n", html)
	return []byte(msg)
}

func main() {
	auth := connect()

	addr := fmt.Sprintf("%s:%d", host, port)
	msg := createMsg("<h1> Heading </h1> 123")

	smtp.SendMail(addr, auth, from, []string{to}, msg)
}
