package server

import (
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/catchsdbot/infra/vars"
)

var botThread sync.Mutex

func StartBot() bool {
	success := botThread.TryLock()
	if success {
		go botMain()
	}
	return success
}

func botMain() {
	defer botThread.Unlock()

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

		commandDispatcher(update.Message)
	}
}

func commandDispatcher(msg *tgbotapi.Message) {
	command := msg.Command()
	args := strings.Fields(msg.CommandArguments())
	logrus.Debugln("Received", command, args)

}
