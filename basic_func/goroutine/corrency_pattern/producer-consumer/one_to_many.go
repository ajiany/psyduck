package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"time"
)

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

var oneToManyChan chan *Pagination
var exitP chan struct{}

func Producer(itemCount int, pageSize int) {
	var pageCount int
	if itemCount != 0 {
		if itemCount%pageSize == 0 {
			pageCount = itemCount / pageSize
		} else {
			pageCount = (itemCount / pageSize) + 1
		}
	}

	for i := 1; i <= pageCount; i++ {
		oneToManyChan <- &Pagination{Page: i, PageSize: pageSize}
	}
	close(exitP)
}

func Consume(idx int) error {
	for {
		select {
		case <-exitP:
			time.Sleep(time.Second)
			fmt.Println(fmt.Sprintf("goroutine %v exitP......", idx))
			return nil
		default:
			stc := <-oneToManyChan
			// 处理业务
			fmt.Println(fmt.Sprintf("goroutine %v woking.... %v", idx, stc))
			//if stc.Page == 49 {
			//	return nil
			//}
		}
	}
	return nil
}

func OneToManyDemo() {
	n, length, segment := 20, 1000, 1000
	oneToManyChan = make(chan *Pagination, n)
	exitP = make(chan struct{}, 0)

	eg, _ := errgroup.WithContext(context.TODO())

	producerFunc := func() error {
		Producer(length, segment)
		return nil
	}

	eg.Go(producerFunc)

	for i := 0; i < n; i++ {
		tmp := i
		eg.Go(func() error {
			return Consume(tmp)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(errors.WithStack(err))
	} else {
		fmt.Println("success")
	}
}
