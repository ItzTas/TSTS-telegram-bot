package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cfg *config) helpCommandCallback(_ tgbotapi.Update, bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, args ...string) error {
	if len(args) > 1 {
		return cfg.helpEspecificCMD(bot, msg, args[1])
	}
	msg.Text += "---Availables commands---\n\n"
	for _, command := range cfg.GetCommands() {
		msg.Text += fmt.Sprintf("Command name: /%v\n", command.commandsInfo.Command)
		msg.Text += fmt.Sprintf("Command description:\n - %v\n\n", command.commandsInfo.Description)
	}
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *config) helpEspecificCMD(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, cmd string) error {
	command, ok := cfg.GetCommands()[cmd]
	if !ok {
		return fmt.Errorf("command %v does not exist", cmd)
	}
	msg.Text += fmt.Sprintf("Command name: /%v\n", command.commandsInfo.Command)
	msg.Text += fmt.Sprintf("Command description:\n - %v\n\n", command.commandsInfo.Description)
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
