package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	Max             = 1000
	manyToManyChan  = make(chan int, 100)
	manyToManyStopC = make(chan struct{}, 0)
)

func ManyToManyProducer(producerNum int, stopC chan struct{}) {

	for i := 0; i < producerNum; i++ {

		go func(exit chan struct{}, idx int) {
			for {
				select {
				case <-exit:
					fmt.Println(fmt.Sprintf("producer %v exit", idx))
					time.Sleep(time.Second)
					return
				default:
					msg := rand.Intn(Max)
					fmt.Println(fmt.Sprintf("producer%v %v", idx, msg))
					manyToManyChan <- msg
					time.Sleep(time.Millisecond * 50)
				}
			}
		}(stopC, i)
	}

}

func ManyToManyConsumer() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		for msg := range manyToManyChan {

			// Todo 业务逻辑
			fmt.Println(fmt.Sprintf("consumer receive msg %v", msg))

			// 退出逻辑
			if msg == Max-1 {
				fmt.Println("send stop signal to producers")
				close(manyToManyStopC)
				return
			}
		}
	}()

	wg.Wait()
}

func ManyToManyDemo() {

	go ManyToManyProducer(10, manyToManyStopC)
	go ManyToManyConsumer()

	time.Sleep(time.Second * 10)
}
