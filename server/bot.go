package server

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/catchsdbot/vars"
)

func Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := vars.BotInstance.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		if update.Message.From.IsBot {
			continue
		}

		logrus.Debugln("Recive command", update.Message.Command())
		// TODO: 在此处理命令
	}
	return nil
}
