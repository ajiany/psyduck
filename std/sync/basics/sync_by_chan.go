package basics

import "time"

// UseChanSyncDemo 使用chan完成同步操作
func UseChanSyncDemo() {

	ch := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		ch <- struct{}{}
		println("done")
	}()
	<-ch

}
