package kmutex

import (
	"hash/fnv"
	"runtime"
	"sync"
)

func NewStringKmutex(n int) IKmutex[string] {
	if n <= 0 {
		n = runtime.NumCPU()
	}
	return &stringKeyMutex{
		mutexes: make([]sync.Mutex, n),
	}
}

type stringKeyMutex struct {
	mutexes []sync.Mutex
}

func (km *stringKeyMutex) Lock(id string) {
	km.mutexes[km.hash(id)%uint32(len(km.mutexes))].Lock()
}

func (km *stringKeyMutex) Unlock(id string) {
	km.mutexes[km.hash(id)%uint32(len(km.mutexes))].Unlock()
}

func (km *stringKeyMutex) hash(id string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(id))
	return h.Sum32()
}
