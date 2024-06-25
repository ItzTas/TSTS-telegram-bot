package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ItzTas/TSTSbot/internal/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	DefaultTimeOut = 5 * time.Minute
)

func startBot(cfg *config) {
	bot, err := tgbotapi.NewBotAPI(cfg.ConfigKeys["botToken"])
	if err != nil {
		fmt.Println("Could not create bot API ", err)
		return
	}

	commands := []tgbotapi.BotCommand{}
	for _, cmd := range cfg.GetCommands() {
		commands = append(commands, cmd.commandsInfo)
	}

	setComandsConfig := tgbotapi.NewSetMyCommands(commands...)
	_, err = bot.Request(setComandsConfig)
	if err != nil {
		fmt.Println("could not set commands ", err)
		return
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.From.ID, "")
		if update.Message.IsCommand() {
			args := strings.Split(update.Message.Text, " ")
			command, exists := cfg.GetCommands()[update.Message.Command()]
			if !exists {
				msg.Text = "Command does not exist"
				_, err := bot.Send(msg)
				if err != nil {
					fmt.Println("Could not send message ", err)
					continue
				}
				continue
			}

			err = command.callback(update, bot, &msg, args...)
			if err != nil {
				msg.Text = err.Error()
				_, err := bot.Send(msg)
				if err != nil {
					fmt.Println("Could not send message ", err)
					continue
				}
			}
			continue
		}

		if err := cfg.getChatResponse(bot, &msg, update.Message.Text); err != nil {
			msg.Text = err.Error()
			_, err := bot.Send(msg)
			if err != nil {
				fmt.Println("Could not send message ", err)
				continue
			}
		}
	}
}

type config struct {
	client     client.Client
	ConfigKeys map[string]string
}

type handlerCommands struct {
	commandsInfo tgbotapi.BotCommand
	callback     func(tgbotapi.Update, *tgbotapi.BotAPI, *tgbotapi.MessageConfig, ...string) error
}

func (cfg *config) GetCommands() map[string]handlerCommands {
	return map[string]handlerCommands{
		"help": {
			commandsInfo: tgbotapi.BotCommand{
				Command:     "help",
				Description: "Displays the availables commands",
			},
			callback: cfg.helpCommandCallback,
		},
		"randomfact": {
			commandsInfo: tgbotapi.BotCommand{
				Command:     "randomfact",
				Description: "Shows a random fact",
			},
			callback: cfg.randomFactCommandCallback,
		},
		"randomcatimage": {
			commandsInfo: tgbotapi.BotCommand{
				Command:     "randomcatimage",
				Description: "shows randoms images of cute cats\n write a number after the comand to define the limit with 12 beeing the maximum",
			},
			callback: cfg.randomCatImageCommandCallback,
		},
	}
}
