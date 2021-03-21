package singleton

import "sync"

type LockCheckLazySingleton struct {
	data int64
}

var ll sync.Mutex
var lcSingleton *LockCheckLazySingleton

func init() {
	ll = sync.Mutex{}
}

func LockCheckLazySingletonInstance() *LockCheckLazySingleton {
	if lcSingleton == nil {
		ll.Lock()
		defer ll.Unlock()
		if lcSingleton == nil {
			lcSingleton = &LockCheckLazySingleton{}
		}
	}
	return lcSingleton
}
