package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not get enviroment variables ", err)
		return
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		fmt.Println("Could not create bot API ", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
}
