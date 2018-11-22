package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/vjeantet/jodaTime"

	"main/api"
	"main/mail"
)

func init() {
	time.LoadLocation("NST") // set timezone to taiwan
}

func sendDailyBeauty(subscribers []string, isTest bool) {
	beauties, err := api.FetchBeauties()

	if err != nil {
		panic(err)
	}

	html := mail.GenerateHTML(beauties)

	date := jodaTime.Format("YYYY-MM-dd", time.Now())
	subject := fmt.Sprintf("表特日報-%s", date)

	if isTest {
		subject += " " + strconv.Itoa(rand.Int())
	}

	for _, to := range subscribers {
		mail.Send(to, subject, html)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	toMails := []string{"pudding850806@gmail.com"}
	sendDailyBeauty(toMails, true)
	log.Println("Test successfully")
	w.Write([]byte("Test successfully"))
}

func publishHandler(w http.ResponseWriter, r *http.Request) {
	toMails := []string{"pudding850806@gmail.com", "w5151381guy@gmail.com", "vorkibiz@gmail.com"}
	sendDailyBeauty(toMails, false)
	log.Println("Publish successfully")
	w.Write([]byte("Publish successfully"))
}

func main() {
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/publish", publishHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listen on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	panic(err)
}
