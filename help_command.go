package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func helpCommandCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	msg.Text += "---Availables commands---\n"
	for _, command := range GetCommands() {
		msg.Text += fmt.Sprintf("Command name: /%v\n", command.commandsInfo.Command)
		msg.Text += fmt.Sprintf("Command description:\n - %v\n", command.commandsInfo.Description)
	}
	return nil
}
