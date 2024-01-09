package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
使用select 从chan中接收数据
*/
func selectDataFromChan() {
	ch1, ch2 := make(chan string, 0), make(chan string, 0)

	go func() {
		for {
			time.Sleep(time.Second)
			ch1 <- fmt.Sprintf("msg1 => %v", rand.Uint32())
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 5)
			ch2 <- fmt.Sprintf("msg2 => %v", rand.Uint32())
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("print data", msg1)
		case msg2 := <-ch2:
			fmt.Println("print data", msg2)
		default:
			break
			fmt.Println("no data print")
		}
	}

}
