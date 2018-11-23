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

	"main/mail"
	"main/ptt"
)

func sendDailyBeauty(subscribers []string, isTest bool) {
	log.Println("getting daily beauty...")
	beauties, err := ptt.FetchBeauties()
	randomBeauty := ptt.FetchRandomBeauty()

	if err != nil {
		panic(err)
	}

	log.Println("generating HTML...")
	html := mail.GenerateHTML(beauties, randomBeauty)

	loc, _ := time.LoadLocation("Asia/Taipei")
	date := jodaTime.Format("YYYY-MM-dd", time.Now().In(loc))
	subject := fmt.Sprintf("表特日報-%s", date)

	if isTest {
		subject += " " + strconv.Itoa(rand.Int())
	}

	log.Println("sending...")
	for _, to := range subscribers {
		mail.Send(to, subject, html)
	}

	log.Println("Finish")
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

// TODO: analysis
// TODO: 禮拜幾標題變化
// TODO: 下載所有圖片
// TODO: 防止手動觸法 cron
