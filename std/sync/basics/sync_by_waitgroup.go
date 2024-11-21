package basics

import (
	"fmt"
	"sync"
)

func UseWaitGroupSyncDemo() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fmt.Println(idx, " ", "done")
		}(i)
	}

	wg.Wait()
}
