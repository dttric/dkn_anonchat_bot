package main

import (
	"log"
	"os"
	"time"

	env "github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func init() {
	if err := env.Load(); err != nil {
		log.Print(".env не найден!")
	}
}

func main() {
	token, _ := os.LookupEnv("TOKEN")
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// надо будет позже
	// b.Handle(tele.OnText, func(c tele.Context) error {
	// 	// All the text messages that weren't
	// 	// captured by existing handlers.

	// 	var (
	// 		user = c.Sender()
	// 		text = c.Text()
	// 	)

	// 	// Use full-fledged bot's functions
	// 	// only if you need a result:
	// 	_, err := b.Send(user, text)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Instead, prefer a context short-hand:
	// 	return c.Send(text)
	// })

	// надо будет позже
	// b.Handle(tele.OnPhoto, func(c tele.Context) error {
	// 	// Photos only.
	// 	photo := c.Message().Photo
	// })

	b.Handle("/help", func(c tele.Context) error {
		return c.Send(`Привет, это бот для анонимного общения.
Все бесплатно и без рекламы.

Для начала зарегистрируйся:
/reg ЛюбойНикнейм
Установка адресата - :НикАдресата

Исходный код: https://github.com/dttric/dkn_anonchat_bot
Связь с админом: :silent`)
	})

	b.Start()
}
