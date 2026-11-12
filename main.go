package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := LoadConfig()

	if cfg.BotToken == "" {
		log.Fatal("BOT_TOKEN belum diatur di .env")
	}

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatalf("Gagal menghubungkan bot: %v", err)
	}

	log.Printf("Bot terhubung: @%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		var response string

		switch update.Message.Command() {
		case "start":
			response = "Selamat datang di Dans AutoTag!\nGunakan /joinalerts untuk mendaftar notifikasi."
		case "joinalerts":
			response = "Pendaftaran notifikasi berhasil diterima."
		default:
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		if _, err := bot.Send(msg); err != nil {
			log.Println("Gagal mengirim pesan:", err)
		}
	}
}
