package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func HelloWorldGoroutine0() {
	funcDemo := func() {
		for i := 0; ; i++ {
			fmt.Println("num => ", i)
		}
	}

	go funcDemo()

	fmt.Println("start...")
	time.Sleep(1 * time.Millisecond)
	fmt.Println("end...")
}

func HelloWoldGoroutine1() {
	funcDemo := func(ch chan string) {
		for {
			ch <- fmt.Sprintf("msg => %v", rand.Uint32())
			time.Sleep(time.Second)
		}
	}
	ch := make(chan string)
	go funcDemo(ch)

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

	time.Sleep(5 * time.Second)
}
