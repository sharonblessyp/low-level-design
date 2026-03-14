package singleton

import "sync"

type threadsafe struct{}

var (
	threadSafeInstance *threadsafe
	mutex              *sync.Mutex
)

func GetThreadSafeInstance() *threadsafe {
	mutex.Lock()
	defer mutex.Unlock()

	if threadSafeInstance == nil {
		threadSafeInstance = &threadsafe{}
	}
	return threadSafeInstance
}
