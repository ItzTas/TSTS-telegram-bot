package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cfg *config) randomFactCommandCallback(_ tgbotapi.Update, bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, args ...string) error {
	fact, err := cfg.client.GetUselessFact()
	if err != nil {
		return fmt.Errorf("could not get random fact: %v", err)
	}
	msg.Text = fact.Text
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
