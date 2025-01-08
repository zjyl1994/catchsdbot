package kmutex

import (
	"runtime"
	"sync"
)

func NewInt64Kmutex(n int) IKmutex[int64] {
	if n <= 0 {
		n = runtime.NumCPU()
	}
	return &int64KeyMutex{
		mutexes: make([]sync.Mutex, n),
	}
}

type int64KeyMutex struct {
	mutexes []sync.Mutex
}

func (km *int64KeyMutex) Lock(id int64) {
	km.mutexes[id%int64(len(km.mutexes))].Lock()
}

func (km *int64KeyMutex) Unlock(id int64) {
	km.mutexes[id%int64(len(km.mutexes))].Unlock()
}
