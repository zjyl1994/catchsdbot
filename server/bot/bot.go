package bot

import (
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/catchsdbot/infra/utils"
	"github.com/zjyl1994/catchsdbot/infra/vars"
	"github.com/zjyl1994/catchsdbot/service/user"
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

	logrus.Infoln("Bot started")

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

		err := commandDispatcher(update.Message)
		if err != nil {
			logrus.Errorln(err)
			utils.ReplyTextToTelegram(update.Message, "发生错误，请联系管理员", false)
		}
	}
}

func commandDispatcher(msg *tgbotapi.Message) error {
	command := msg.Command()
	args := strings.Fields(msg.CommandArguments())
	logrus.Debugln("Received", command, args)

	switch command {
	case "getsp":
		return handleGetSP(msg)
	case "catchsd":
		return handleCatchSd(msg)
	default:
		utils.ReplyTextToTelegram(msg, "未知命令", false)
		return nil
	}
}

func getUser(msg *tgbotapi.Message) (*user.User, error) {
	tgUserId := msg.From.ID
	tgUserName := msg.From.FirstName + " " + msg.From.LastName
	return user.GetOrCreateByTgUser(tgUserId, tgUserName)
}
