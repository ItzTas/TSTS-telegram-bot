package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ItzTas/TSTSbot/internal/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

const (
	DefaultTimeOut = 5 * time.Minute
)

func startBot(cfg *config) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not get enviroment variables ", err)
		return
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		fmt.Println("Could not create bot API ", err)
		return
	}

	bot.Debug = true

	commands := []tgbotapi.BotCommand{}
	for _, cmd := range GetCommands(cfg) {
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
			err := GetCommands(cfg)[update.Message.Command()].callback(bot, &update, &msg)
			if err != nil {
				msg.Text = err.Error()
			}
		} else {
			msg.Text = update.Message.Text
		}
		_, err := bot.Send(msg)
		if err != nil {
			fmt.Println("Could not send message ", err)
			continue
		}
	}
}

type config struct {
	client client.Client
}

type handlerCommands struct {
	commandsInfo tgbotapi.BotCommand
	callback     func(*tgbotapi.BotAPI, *tgbotapi.Update, *tgbotapi.MessageConfig) error
}

func GetCommands(cfg *config) map[string]handlerCommands {
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
	}
}
