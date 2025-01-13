package controller

import (
	"github.com/zjyl1994/catchsdbot/infra/vars"
	"github.com/zjyl1994/catchsdbot/service/cargo"
	"github.com/zjyl1994/catchsdbot/service/dice"
	"github.com/zjyl1994/catchsdbot/service/stamina"
	"gorm.io/gorm"
)

const CATCHSD_SP = 10 // 每次捕捉需要10体力值

type CatchSdResult struct {
	Result  int
	Dice    int
	Message string
	Amount  int64
}

var (
	catchsdMessage = []string{
		"深渊迷航，你的企鹅在路上折损大半",
		"一波三折，尽管有些企鹅中途逃离，但你还是成功回来了",
		"风平浪静，你成功的带回了属于你的企鹅",
		"一帆风顺，路上有些企鹅加入了你的队伍",
		"好运相随，幸运的人总是能获得更多企鹅",
	}
)

func CatchSd(userId int64, num int) (*CatchSdResult, error) {
	catchSP := num * CATCHSD_SP
	_, err := stamina.UseStaminPoint(userId, int64(catchSP))
	if err != nil {
		return nil, err
	}
	result := new(CatchSdResult)
	// 扔骰子决定本次的结果
	result.Dice = dice.Roll()
	result.Result = int((1 + dice.GetDiceBuff(result.Dice)) * float64(num))
	result.Message = dice.GetByDiceResult(result.Dice, catchsdMessage)
	// 写入库存
	err = vars.Database.Transaction(func(tx *gorm.DB) error {
		amount, err := cargo.GetCargoItem(tx, userId, cargo.ITEM_PENGUIN)
		if err != nil {
			return err
		}
		result.Amount = amount + int64(result.Result)
		return cargo.SetCargoItem(tx, userId, cargo.ITEM_PENGUIN, result.Amount)
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
