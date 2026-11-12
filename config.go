package main

import (
	"os"
	"strings"
)

type Config struct {
	BotToken        string
	ForceSubChannel string
}

func LoadConfig() Config {
	return Config{
		BotToken:        strings.TrimSpace(os.Getenv("BOT_TOKEN")),
		ForceSubChannel: strings.TrimSpace(os.Getenv("FORCE_SUB_CHANNEL")),
	}
}
