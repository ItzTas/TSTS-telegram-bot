package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cfg *config) randomFactCommandCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	fact, err := cfg.client.GetUselessFact()
	if err != nil {
		return fmt.Errorf("could not get random fact: %v", err)
	}
	msg.Text = fact.Text
	return nil
}
