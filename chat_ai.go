package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (cfg *config) getChatResponse(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, content string) error {
	res, err := cfg.client.GetaiChat(content)
	if err != nil {
		return err
	}

	msg.Text = res
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
