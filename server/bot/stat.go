package bot

import (
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zjyl1994/catchsdbot/infra/utils"
	"github.com/zjyl1994/catchsdbot/service/cargo"
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

func handleMyStat(msg *tgbotapi.Message) error {
	u, err := getUser(msg)
	if err != nil {
		return err
	}

	var sb strings.Builder
	sb.WriteString("玩家：*" + utils.EscapeTelegramMarkdown(u.TgUserName) + "*\n\n")

	sp, err := stamina.GetStaminPoint(u.ID)
	if err != nil {
		return err
	}
	sb.WriteString(sp.String() + "\n")
	sb.WriteString("\n*库存信息*：\n")
	cargoInfo, err := cargo.GetCargo(u.ID)
	if err != nil {
		return err
	}
	for ctype, amount := range cargoInfo {
		if name, ok := cargo.ItemName[ctype]; ok {
			sb.WriteString(name)
			sb.WriteString(":")
			sb.WriteString(strconv.FormatInt(amount, 10))
			sb.WriteString("\n")
		}
	}
	return utils.ReplyTextToTelegram(msg, sb.String(), true)
}
