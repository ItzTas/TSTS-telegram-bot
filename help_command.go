package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cfg *config) helpCommandCallback(bot *tgbotapi.BotAPI, _ *tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	msg.Text += "---Availables commands---\n\n"
	for _, command := range GetCommands(cfg) {
		msg.Text += fmt.Sprintf("Command name: /%v\n", command.commandsInfo.Command)
		msg.Text += fmt.Sprintf("Command description:\n - %v\n\n", command.commandsInfo.Description)
	}
	return nil
}
