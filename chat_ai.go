package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (cfg *config) getChatResponse(msg *tgbotapi.MessageConfig, content string) error {
	res, err := cfg.client.GetaiChat(content)
	if err != nil {
		return err
	}

	msg.Text = res
	return nil
}
