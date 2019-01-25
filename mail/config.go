package mail

import "os"

var (
	host = os.Getenv("SMTP_HOST")
	port = os.Getenv("SMTP_PORT")
	user = os.Getenv("SMTP_USER")
	pwd  = os.Getenv("SMTP_PWD")
)
