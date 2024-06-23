package main

import (
	"fmt"
	"os"

	"github.com/ItzTas/TSTSbot/internal/client"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not get enviroment variables ", err)
		return
	}

	catAPIKey := os.Getenv("CAT_API_KEY")
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	cfg := config{
		ConfigKeys: map[string]string{
			"botToken": botToken,
		},
		client: client.NewClient(DefaultTimeOut, map[string]string{
			"catAPIKey": catAPIKey,
		}),
	}

	startBot(&cfg)
}
