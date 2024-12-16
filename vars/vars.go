package vars

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	DebugMode bool

	BotToken    string
	BotInstance *tgbotapi.BotAPI
)
