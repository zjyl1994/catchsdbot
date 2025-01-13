package bot

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zjyl1994/catchsdbot/controller"
	"github.com/zjyl1994/catchsdbot/infra/utils"
	"github.com/zjyl1994/catchsdbot/service/stamina"
)

func handleCatchSd(msg *tgbotapi.Message) error {
	u, err := getUser(msg)
	if err != nil {
		return err
	}
	args := msg.CommandArguments()
	var catchNum int
	if strings.ToLower(args) == "all" {
		sp, err := stamina.GetStaminPoint(u.ID)
		if err != nil {
			return err
		}
		catchNum = int(sp.Current() / controller.CATCHSD_SP)
	} else if args == "" {
		catchNum = 1
	} else {
		catchNum, err = strconv.Atoi(args)
		if err != nil {
			return utils.ReplyTextToTelegram(msg, "无法解析 "+args, false)
		}
	}
	result, err := controller.CatchSd(u.ID, catchNum)
	if err != nil {
		return err
	}
	resultMsg := fmt.Sprintf("**ROLL**\\(6\\)\\=%d\n%s\n本次共捕捉到%d个企鹅\n\n现有鹅口%d", result.Dice, result.Message, result.Result, result.Amount)
	return utils.ReplyTextToTelegram(msg, resultMsg, true)
}
