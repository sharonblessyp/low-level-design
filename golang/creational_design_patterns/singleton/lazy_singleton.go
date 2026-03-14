package singleton

import "sync"

type LazySingleton struct {
}

var (
	lazyinstance *LazySingleton
	once         *sync.Once
)

func GetLazyInstance() *LazySingleton {
	once.Do(func() {
		lazyinstance = &LazySingleton{}
	})
	return lazyinstance
}
