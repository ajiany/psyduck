package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var manyToOneChan = make(chan string, 20)

func ManyToOneProducer(num int) {
	wg := sync.WaitGroup{}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 500)
			msg := fmt.Sprintf("idx:%v:%v", idx, rand.Uint32())
			fmt.Println("producer send msg => ", msg)
			manyToOneChan <- msg
		}(i)
	}

	wg.Wait()
}

func ManyToOneConsumer() {
	for {
		select {
		case msg := <-manyToOneChan:
			fmt.Println(fmt.Sprintf("consumer receive msg => %v", msg))
		}
	}
}

func ManyToOneDemo() {
	go ManyToOneProducer(10)
	go ManyToOneConsumer()

	time.Sleep(5 * time.Second)
}
