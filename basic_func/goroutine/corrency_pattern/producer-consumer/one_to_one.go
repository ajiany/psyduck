package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
生产者-消费者模式
*/

var one2OneChan = make(chan string, 5)

//
// One2OneProducer
//  @Description: 生产者与消费者1:1 模拟生产者产生数据
//  @param wg
//  @param producerStopC
//
func One2OneProducer(wg *sync.WaitGroup, producerStopC chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-producerStopC:
			fmt.Println("producer exit...")
			return
		default:
			fmt.Println(fmt.Sprintf("=> chan  %v", rand.Uint32()))
			one2OneChan <- fmt.Sprintf("%v", rand.Uint32())
			time.Sleep(time.Millisecond * 200)
		}
	}
}

//
// One2OneConsumer
//  @Description: 生产者与消费者1:1 模拟消费者产生数据
//  @param wg
//  @param consumerStopC
//
func One2OneConsumer(wg *sync.WaitGroup, consumerStopC chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-consumerStopC:
			fmt.Println("consumer exit...")
			return
		case msg := <-one2OneChan:
			fmt.Println(fmt.Sprintf("chan => %v", msg))
		default:
			time.Sleep(time.Second)
		}
	}
}

func One2OneDemo() {
	producerStopC, consumerStopC := make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	go One2OneProducer(&wg, producerStopC)

	go One2OneConsumer(&wg, consumerStopC)

	// 等待5秒再让生产者结束
	time.Sleep(time.Second * 5)
	producerStopC <- struct{}{}

	// 等待5秒再让消费者结束
	time.Sleep(time.Second * 2)
	close(one2OneChan)
	consumerStopC <- struct{}{}

	wg.Wait()
}
