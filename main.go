package main

import (
	"github.com/ItzTas/TSTSbot/internal/client"
)

func main() {

	cfg := config{
		client: client.NewClient(DefaultTimeOut),
	}
	startBot(&cfg)
}
