package vars

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

var (
	DebugMode bool

	DataDir    string
	ListenAddr string

	BotToken    string
	BotInstance *tgbotapi.BotAPI

	Database *gorm.DB
)
