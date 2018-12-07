package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"main/model"
)

const (
	name = "Daily Beauty"
	from = "service@daily-beauty.xyz"
)

var auth = smtp.PlainAuth("", user, pwd, host)

func createMsg(to, subject, html string) []byte {
	msg := ""
	msg += fmt.Sprintf("To: %s\r\n", to)
	msg += fmt.Sprintf("From: %s <%s>\r\n", name, from)
	msg += fmt.Sprintf("Subject: %s\r\n", subject)

	msg += "MIME-version: 1.0;\r\n"
	msg += `Content-Type: text/html; charset="UTF-8"` + "\r\n"

	msg += "\r\n"
	msg += fmt.Sprintf("%s\r\n", html)
	return []byte(msg)
}

// Send sends the html to the receiver
func Send(to, subject, html string) {
	addr := fmt.Sprintf("%s:%d", host, port)
	msg := createMsg(to, subject, html)
	smtp.SendMail(addr, auth, from, []string{to}, msg)
}

func reverse(bs []model.Beauty) []model.Beauty {
	n := len(bs)
	reversedBs := make([]model.Beauty, n)
	for i, b := range bs {
		reversedBs[n-i-1] = b
	}
	return reversedBs
}

// GenerateHTML generates a html from beauties slice
func GenerateHTML(beauties []model.Beauty, randomBeauty model.Beauty, token string) string {
	tmpl := template.Must(template.ParseFiles("./mail/mail.html"))

	data := map[string]interface{}{
		"Beauties":     reverse(beauties),
		"RandomBeauty": randomBeauty,
		"Token":        token,
	}

	var b bytes.Buffer
	tmpl.Execute(&b, data)
	html := b.String()

	return html
}
