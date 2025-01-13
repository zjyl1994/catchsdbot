package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zjyl1994/catchsdbot/infra/utils"
	"github.com/zjyl1994/catchsdbot/service/stamina"
)

func handleGetSP(msg *tgbotapi.Message) error {
	u, err := getUser(msg)
	if err != nil {
		return err
	}
	sp, err := stamina.GetStaminPoint(u.ID)
	if err != nil {
		return err
	}
	return utils.ReplyTextToTelegram(msg, sp.String(), false)
}
