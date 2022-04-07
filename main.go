package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func captchaMessageFactory(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter a+b=c")

	btn := tgbotapi.NewInlineKeyboardButtonData("1", "1")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("2", "2")

	markup := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(btn, btn2))

	msg.ReplyMarkup = markup
	msg.DisableNotification = true

	return msg
}

func getBanDuration(seconds int64) int64 {
	return time.Now().Add(time.Duration(seconds) * time.Second).Unix()
}

func main() {
	bot, err := tgbotapi.NewBotAPI("5193948128:AAE7L1OEzT6J1iG3yNAQ-VKHwSPtoDrZ66M")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		fmt.Println("up")

		fmt.Printf("CHAT [%v] ", update.Message.NewChatMembers)

		if update.Message == nil {

			continue
		}

		kick := tgbotapi.KickChatMemberConfig{
			ChatMemberConfig: tgbotapi.ChatMemberConfig{
				ChatID: update.Message.Chat.ID,
				UserID: update.Message.From.ID,
			},
			UntilDate: getBanDuration(1),
		}

		bot.Send(kick)
	}

}
