package stamina

import (
	"errors"
	"sync"
)

const (
	SP_MAX        = 200 // 体力自然恢复上限
	SP_PER_SECOND = 400 // 400秒一点体力
)

var (
	globalSPLock sync.Mutex //全局体力锁，后期性能不佳时改为用户级锁
	ErrNotEnough = errors.New("not enough stamina point")
)
