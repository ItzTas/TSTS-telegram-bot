package main

import (
	"fmt"

	"github.com/ItzTas/TSTSbot/internal/client"
)

func main() {

	cfg := config{
		client: client.NewClient(DefaultTimeOut),
	}
	fmt.Println(cfg.client.GetaiChat("who are you?"))
	startBot(&cfg)
}
