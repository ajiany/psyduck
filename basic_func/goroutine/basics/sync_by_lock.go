package basics

import "sync"

func UseLockSyncDemo() {
	var lock sync.Mutex
	lock.Lock()
	go func() {
		lock.Unlock()
		println("done")
	}()

	lock.Lock()
	lock.Unlock()
	println("exit")
}
