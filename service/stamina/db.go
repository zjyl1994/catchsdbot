package stamina

import (
	"errors"
	"fmt"
	"time"

	"github.com/zjyl1994/catchsdbot/infra/utils"
	"github.com/zjyl1994/catchsdbot/infra/vars"
	"gorm.io/gorm"
)

func GetStaminPoint(userId int64) (*Stamina, error) {
	var sp Stamina
	err := vars.Database.Where(Stamina{UserId: userId}).First(&sp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 不存在记录的用户直接返回最高能量上限
			sp.UserId = userId
			sp.LastSP = SP_MAX
			sp.LastTick = time.Now().Unix()
			return &sp, nil
		}
		return nil, err
	}
	return &sp, nil
}

func UseStaminPoint(userId int64, cost int64) (*Stamina, error) {
	spLock.Lock(userId)
	defer spLock.Unlock(userId)

	sp, err := GetStaminPoint(userId)
	if err != nil {
		return nil, err
	}
	// 计算并扣减能量
	current := sp.Current()
	remainEnergy := current - cost
	// 检查是否扣完
	if remainEnergy < 0 {
		return sp, utils.NewBizErr(fmt.Sprintf("SP不足%d,%s", cost, sp.String()), ErrNotEnough)
	}
	// 新体力写入DB
	sp.LastSP = remainEnergy
	sp.LastTick = time.Now().Unix()

	err = vars.Database.Save(sp).Error
	if err != nil {
		return nil, err
	}
	return sp, nil
}

func AddStaminPoint(userId int64, amount int64) error {
	spLock.Lock(userId)
	defer spLock.Unlock(userId)

	sp, err := GetStaminPoint(userId)
	if err != nil {
		return err
	}

	sp.LastSP = sp.Current() + amount
	sp.LastTick = time.Now().Unix()

	return vars.Database.Save(sp).Error
}
