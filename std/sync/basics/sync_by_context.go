package basics

import (
	"context"
	"time"
)

func UseContextSyncDemo() {
	ctx, cancel := context.WithCancel(context.TODO())
	go func() {
		defer cancel()
		time.Sleep(5 * time.Second)
		println("done")
	}()

	<-ctx.Done()
}
