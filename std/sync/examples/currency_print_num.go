package main

import (
	"fmt"
)

// currencyPrintNum 使用n个goroutine打印m个数字
func currencyPrintNum() {
	number, goCount := 100, 5
	exitCh := make(chan struct{}, 0)
	ch := make(chan struct{}, goCount)

	for i := 0; i < number; i++ {
		ch <- struct{}{}
		go func(num int) {
			fmt.Println("number => ", num)
			if num == number-1 {
				exitCh <- struct{}{}
			}
			<-ch
		}(i)

	}

	<-exitCh
}
