package controller

import "github.com/zjyl1994/catchsdbot/service/stamina"

const CATCHSD_SP = 10 // 每次捕捉需要10体力值

type CatchSdResult struct {
	Result []int
}

func CatchSd(n int, userId int64) (*CatchSdResult, error) {
	catchSP := n * CATCHSD_SP
	_, err := stamina.UseStaminPoint(userId, int64(catchSP))
	if err != nil {
		return nil, err
	}
	// TODO：实现捕捉逻辑
	return nil, nil
}
