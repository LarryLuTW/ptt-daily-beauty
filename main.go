package main

import (
	"fmt"
	"log"
	"main/db"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vjeantet/jodaTime"

	"main/mail"
	"main/ptt"
)

func sendDailyBeauty(subscribers []string, isTest bool) {
	log.Println("getting daily beauty...")
	// TODO: do parallelly
	beauties, err := ptt.FetchBeauties()
	if err != nil {
		panic(err)
	}

	randomBeauty, err := ptt.FetchRandomBeauty()
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
		log.Printf("Send to '%s' success", to)
		time.Sleep(200 * time.Millisecond)
	}

	log.Println("Finish")
}

func testHandler(c *gin.Context) {
	toMails := []string{"pudding850806@gmail.com"}
	sendDailyBeauty(toMails, true)
	log.Println("Test successfully")
	c.String(200, "Test successfully")
}

func publishHandler(c *gin.Context) {
	toMails, err := db.GetEmails()
	if err != nil {
		panic(err)
	}

	sendDailyBeauty(toMails, false)
	log.Println("Publish successfully")
	c.String(200, "Publish successfully")
}

// func subscribeHandler(w http.ResponseWriter, r *http.Request) {
// }

// func unsubscribeHandler(w http.ResponseWriter, r *http.Request) {

// }

func main() {
	r := gin.Default()
	r.GET("/test", testHandler)
	r.GET("/publish", publishHandler)

	// http.HandleFunc("/api/subscribe", subscribeHandler)
	// http.HandleFunc("/api/unsubscribe", unsubscribeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listen on port %s", port)
	err := r.Run(":" + port)
	panic(err)
}

// TODO: unit test
// TODO: logging
// TODO: analysis 轉網址
// TODO: 禮拜幾標題變化
// TODO: 下載所有圖片
// TODO: 防止手動觸發 cron
