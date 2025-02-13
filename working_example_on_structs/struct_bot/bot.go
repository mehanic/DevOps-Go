package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	BotToken   = "838000255:AAGl98I4HGRhNwMa1EBLOOqT6r2C6fU9xas"
	WebhookURL = "https://680b65f41cc4.ngrok.io"
)

var rss = map[string]string{
	"Habr": "https://habrahabr.ru/rss/best/",
}

type RSS struct {
	Items []Item `xml:"channel>item"`
}

type Item struct {
	URL   string `xml:"guid"`
	Title string `xml:"title"`
}

func getNews(url string) (*RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rss := new(RSS)
	err = xml.Unmarshal(body, rss)
	if err != nil {
		return nil, err
	}

	return rss, nil
}

func main() {
	// Initialize bot
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize bot: %v", err))
	}

	// Debugging: print bot details
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	// Set webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		panic(fmt.Sprintf("Failed to set webhook: %v", err))
	}

	// Start listening for updates via webhook
	updates := bot.ListenForWebhook("/")

	// Start the server to handle webhook requests
	go func() {
		fmt.Println("Starting server on port 8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(fmt.Sprintf("Failed to start HTTP server: %v", err))
		}
	}()

	// Handling incoming updates
	for update := range updates {
		if url, ok := rss[update.Message.Text]; ok {
			rss, err := getNews(url)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, an error occurred while fetching the news."))
				continue
			}
			for _, item := range rss.Items {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n%s", item.URL, item.Title)))
			}
		} else {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "There is only the Habr feed available"))
		}
	}
}


// Authorized on account RZCbot
// Starting server on port 8080...
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x50 pc=0x6b9202]

// goroutine 1 [running]:
// main.main()
