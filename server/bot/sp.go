package bot

import (
	"fmt"

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
	current := sp.Current()
	remainSec := sp.RemainSecond()
	var msgText string
	if remainSec > 0 {
		msgText = fmt.Sprintf("当前体力为：%d,距离恢复下一点还有%d秒。", current, remainSec)
	} else {
		msgText = fmt.Sprintf("当前体力为：%d,已达到自然恢复上限。", current)
	}
	return utils.ReplyTextToTelegram(msg, msgText, false)
}
