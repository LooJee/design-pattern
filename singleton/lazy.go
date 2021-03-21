package singleton

import (
	"sync"
	"sync/atomic"
)

type LazySingleton struct {
	data uint64
}

func (l *LazySingleton) Incr() uint64 {
	return atomic.AddUint64(&l.data, 1)
}

var lazySingleton *LazySingleton
var lock sync.Mutex

func init() {
	lock = sync.Mutex{}
}

func LazySingletonInstance() *LazySingleton {
	lock.Lock()
	defer lock.Unlock()
	if lazySingleton == nil {
		lazySingleton = &LazySingleton{}
	}
	return lazySingleton
}

