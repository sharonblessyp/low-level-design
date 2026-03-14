package singleton

import "sync"

type dCThreadSafe struct{}

var (
	dCdCThreadSafeInstance *dCThreadSafe
	dcMutex                *sync.Mutex
)

func GetdCdCThreadSafeInstance() *dCThreadSafe {
	if dCdCThreadSafeInstance == nil {
		dcMutex.Lock()
		defer dcMutex.Unlock()

		if dCdCThreadSafeInstance == nil {
			dCdCThreadSafeInstance = &dCThreadSafe{}
		}
	}
	return dCdCThreadSafeInstance
}
