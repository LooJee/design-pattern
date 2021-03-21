package singleton

import (
	"sync"
	"sync/atomic"
)

type CheckLazySingleton struct {
	data uint64
}

func (l *CheckLazySingleton) Incr() uint64 {
	return atomic.AddUint64(&l.data, 1)
}

var checkLazySingleton *CheckLazySingleton
var checkOnce = sync.Once{}

func CheckLazySingletonInstance() *CheckLazySingleton {
	if checkLazySingleton == nil {
		checkOnce.Do(func() {
			checkLazySingleton = &CheckLazySingleton{}
		})
	}

	return checkLazySingleton
}
