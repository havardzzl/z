package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/havardzzl/havardzzl/apis"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UrlRes struct {
	Url string `json:"url"`
}

func getCatUrl() string {
	catReq, _ := http.NewRequest(http.MethodGet, "https://api.thecatapi.com/v1/images/search", nil)
	catReq.Header.Add("x-api-key", "live_tkFwcVozxR3MvERDQ5isOPNODoBSCNvvQgtU6ckPGiCkqzFYpJTEKV5OsKmXrx4O")
	resp, _ := http.DefaultClient.Do(catReq)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		bs, _ := io.ReadAll(resp.Body)
		ress := []UrlRes{}
		json.Unmarshal(bs, &ress)
		if len(ress) > 0 {
			return ress[0].Url
		}
	}
	return ""
}

type SearchRes struct {
	Value []UrlRes `json:"value"`
}

func searchAndGetRandomUrl(s string) string {
	bs := apis.SearchPic(s)
	res := SearchRes{}
	json.Unmarshal(bs, &res)
	if len(res.Value) > 0 {
		pos := rand.Intn(len(res.Value))
		return res.Value[pos].Url
	}
	return ""
}

func main() {
	bot, err := tgbotapi.NewBotAPI("5878974107:AAGU-k6aPUdy8Sg96H38eKPs7hG6798YSEc")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates := bot.GetUpdatesChan(u)

	// picF, err := os.ReadFile("/Users/zhangzhongliang/Downloads/ttt.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// client := http.Client{
	// 	CheckRedirect: func(r *http.Request, via []*http.Request) error {
	// 		r.URL.Opaque = r.URL.Path
	// 		return nil
	// 	},
	// }

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			var msg tgbotapi.MessageConfig

			// msg.ReplyToMessageID = update.Message.MessageID

			var picUrl string
			searchUrl := searchAndGetRandomUrl(update.Message.Text)
			if searchUrl != "" {
				picUrl = searchUrl
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Here is photo for:"+update.Message.Text)
			} else {
				catUrl := getCatUrl()
				picUrl = catUrl
				if catUrl != "" {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("No photo found for `%s`, here is a cat (:)", update.Message.Text))
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "No photo found for anything!!!")
				}
			}

			bot.Send(msg)

			if len(picUrl) > 0 {
				// photoBs := tgbotapi.FileBytes{
				// 	Name:  "picture",
				// 	Bytes: picF,
				// }
				// photoMsg := tgbotapi.NewPhoto(update.Message.Chat.ID, photoBs)
				photoMsg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(picUrl))
				bot.Send(photoMsg)
			}

		}
	}
}
