package singleton_test

import (
	"design-pattern/singleton"
	"sync"
	"testing"
)

func TestHungrySingletonInstance(t *testing.T) {
	instance := singleton.HungrySingletonInstance()

	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(j int) {
			defer wg.Done()
			tmpInstance := singleton.HungrySingletonInstance()
			if tmpInstance != instance {
				t.Errorf("different instance")
			}
		}(i)
	}

	wg.Wait()
}

func TestLazySingletonInstance(t *testing.T) {
	instance := singleton.LazySingletonInstance()

	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(j int) {
			defer wg.Done()
			tmpInstance := singleton.LazySingletonInstance()
			if tmpInstance != instance {
				t.Errorf("different instance")
			}
		}(i)
	}
	wg.Wait()
}

func BenchmarkHungrySingletonInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.HungrySingletonInstance() != singleton.HungrySingletonInstance() {
				b.Errorf("different instance")
			}
		}
	})
}

func BenchmarkLazySingletonInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.LazySingletonInstance() != singleton.LazySingletonInstance() {
				b.Errorf("different instance")
			}
		}
	})
}

func BenchmarkLockCheckLazySingletonInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.LockCheckLazySingletonInstance() != singleton.LockCheckLazySingletonInstance() {
				b.Errorf("different instance")
			}
		}
	})
}

func BenchmarkCheckLazySingletonInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.CheckLazySingletonInstance() != singleton.CheckLazySingletonInstance() {
				b.Errorf("different instance")
			}
		}
	})
}
