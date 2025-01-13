package bot

import (
	"math/rand/v2"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zjyl1994/catchsdbot/controller"
	"github.com/zjyl1994/catchsdbot/infra/utils"
	"github.com/zjyl1994/catchsdbot/service/stamina"
)

const (
	CATCH_STICKER_SILENT_DEPTH            = "CAACAgEAAxkBAAESUHliQRGsqHOJyBoUehSNURi8UHEEjAACywEAAsRbEEagPE8TOuXrYSME"
	CATCH_STICKER_SILENT_DEPTH_WATCHING_1 = "CAACAgEAAxkBAAES_DJiWTSZkAtJm4jEzXI5EBjOCdywpgACKwMAAmayyEadFj8SUndZ5SME"
	CATCH_STICKER_SILENT_DEPTH_WATCHING_2 = "CAACAgEAAxkBAAEUuKZinapHnTE_rQqyD-vGURecoMv-7wACdQIAArwW8URp5nideF3JuSQE"
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
		if catchNum == 0 {
			return utils.ReplyTextToTelegram(msg, sp.String(), false)
		}
	} else if args == "" {
		catchNum = 1
	} else {
		catchNum, err = strconv.Atoi(args)
		if err != nil || catchNum <= 0 {
			return utils.ReplyTextToTelegram(msg, "无法解析 "+args, false)
		}
	}
	return handleCatchSdWithNum(msg, u.ID, catchNum)
}

func handleCatchSdAll(msg *tgbotapi.Message) error {
	u, err := getUser(msg)
	if err != nil {
		return err
	}
	sp, err := stamina.GetStaminPoint(u.ID)
	if err != nil {
		return err
	}
	catchNum := int(sp.Current() / controller.CATCHSD_SP)
	if catchNum == 0 {
		return utils.ReplyTextToTelegram(msg, sp.String(), false)
	}
	return handleCatchSdWithNum(msg, u.ID, catchNum)
}

func handleCatchSd10(msg *tgbotapi.Message) error {
	u, err := getUser(msg)
	if err != nil {
		return err
	}
	return handleCatchSdWithNum(msg, u.ID, 10)
}

func handleCatchSdWithNum(msg *tgbotapi.Message, userId int64, catchNum int) error {
	result, err := controller.CatchSd(userId, catchNum)
	if err != nil {
		return err
	}
	if catchNum == 1 {
		if result.Result == 0 {
			switch rand.IntN(3) {
			case 0:
				return utils.ReplyTextToTelegram(msg, "手滑了企鹅逃走了", false)
			case 1:
				return utils.ReplyStickerToTelegram(msg, CATCH_STICKER_SILENT_DEPTH_WATCHING_1)
			case 2:
				return utils.ReplyStickerToTelegram(msg, CATCH_STICKER_SILENT_DEPTH_WATCHING_2)
			}
		} else {
			return utils.ReplyStickerToTelegram(msg, CATCH_STICKER_SILENT_DEPTH)
		}
	}
	sp, err := stamina.GetStaminPoint(userId)
	if err != nil {
		return err
	}
	var sb strings.Builder
	sb.WriteString("本次共捕捉到")
	sb.WriteString(strconv.Itoa(result.Result))
	sb.WriteString("个企鹅\n\n")
	sb.WriteString("*ROLL*")
	sb.WriteString(utils.EscapeTelegramMarkdown("(100)="))
	sb.WriteString(strconv.Itoa(result.Dice))
	sb.WriteString("\n")
	sb.WriteString(result.Message)
	sb.WriteString("\n\n")
	sb.WriteString(sp.String())
	return utils.ReplyTextToTelegram(msg, sb.String(), true)
}
