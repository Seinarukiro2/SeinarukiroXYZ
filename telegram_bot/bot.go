package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func startCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я телеграм-бот на Go!")
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	token := "6106150594:AAGWY63P68Bt1RSnFiOhLOtCGRJnS_QUGJA"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					startCommand(bot, update)
				}
			}
		}
	}
}
