package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cfg *config) randomCatImageCommandCallback(update tgbotapi.Update, bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, args ...string) error {
	limit := "1"
	if len(args) > 1 {
		limit = args[1]
	}
	imgs, err := cfg.client.GetCatPictures(limit)
	if err != nil {
		return err
	}

	for _, img := range imgs {
		photo := tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FileURL(img.URL))
		_, err := bot.Send(photo)
		if err != nil {
			fmt.Println("could not send photo", err)
			return err
		}
	}
	msg.Text = "Here are you cat images! :)"
	_, err = bot.Send(msg)
	if err != nil {
		fmt.Println("could not send msg", err)
		return err
	}
	return nil
}
