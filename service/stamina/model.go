package stamina

import (
	"time"

	"github.com/zjyl1994/catchsdbot/infra/utils"
)

type Stamina struct {
	ID       int64 `gorm:"primaryKey"`
	UserId   int64 `gorm:"uniqueIndex,column:user_id"`
	LastTick int64
	LastSP   int64
}

// 计算当前体力
func (s *Stamina) Current() int64 {
	return utils.IdleCalcWithMax(s.LastTick, time.Now().Unix(), s.LastSP, SP_PER_SECOND, SP_MAX)
}

// 计算恢复下一点体力的剩余秒数
func (s *Stamina) RemainSecond() int64 {
	current := s.Current()
	if current >= SP_MAX { // 体力满不会自动回复
		return 0
	}
	elapsedSecond := time.Now().Unix() - s.LastTick
	return SP_PER_SECOND - (elapsedSecond % SP_PER_SECOND)
}
