package vars

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	DebugMode bool

	ListenAddr string

	BotToken    string
	BotInstance *tgbotapi.BotAPI
)
