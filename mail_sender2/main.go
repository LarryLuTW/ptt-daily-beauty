package main

import (
	// "main/mail"

	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("port %s", port)
	http.ListenAndServe(":"+port, nil)

	// resp.Body

	// to := "pudding850806@gmail.com"
	// subject := "This is subject"
	// html := "<h1> Heading 123 </h1> 123"
	// mail.Send(to, subject, html)
}
