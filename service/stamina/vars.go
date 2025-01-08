package stamina

import (
	"errors"

	"github.com/zjyl1994/catchsdbot/infra/utils/kmutex"
)

const (
	SP_MAX        = 200 // 体力自然恢复上限
	SP_PER_SECOND = 400 // 400秒一点体力
)

var (
	spLock       = kmutex.NewInt64Kmutex(0) //体力锁
	ErrNotEnough = errors.New("体力不足")
)
