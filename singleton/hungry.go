package singleton

import "sync/atomic"

type HungrySingleton struct {
	data uint64
}

/*
 add 1 to data, and return the new data
*/
func (s *HungrySingleton) Incr() uint64 {
	return atomic.AddUint64(&s.data, 1)
}

var hungry = &HungrySingleton{}

func HungrySingletonInstance() *HungrySingleton {
	return hungry
}
