package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"time"
)

var (
	Max            = 1000
	manyToManyChan = make(chan *Pagination, 100)
	exitP          = make(chan struct{}, 0)
)

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func Producer(pageSize, pageCount int) {

	for i := 1; i <= pageCount; i++ {
		manyToManyChan <- &Pagination{Page: i, PageSize: pageSize}
	}
	time.Sleep(time.Second * 2)
	close(exitP)
	close(manyToManyChan)
}

func Consumer(idx int) error {

	for {
		select {
		case <-exitP:
			time.Sleep(time.Second)
			fmt.Println(fmt.Sprintf("goroutine %v exitP......", idx))
			return nil
		default:
			stc, ok := <-manyToManyChan
			if !ok {
				return nil
			}
			// 处理业务
			fmt.Println(fmt.Sprintf("goroutine %v woking.... %v", idx, stc))
		}
	}
	return nil

	//eg, _ := errgroup.WithContext(context.TODO())
	//n := 10
	//eg.SetLimit(n)
	//
	//eg.Go(func() error {
	//	for msg := range manyToManyChan {
	//
	//		// Todo 业务逻辑
	//		fmt.Println(fmt.Sprintf("goroutine woking.... %v", msg))
	//		// 退出逻辑
	//		if msg.Page == pageCount {
	//			fmt.Println("send stop signal to producers")
	//			close(manyToManyStopC)
	//			return nil
	//		}
	//	}
	//	return nil
	//})
	//
	//if err := eg.Wait(); err != nil {
	//	fmt.Println(errors.WithStack(err))
	//} else {
	//	fmt.Println("success")
	//}
}

func ManyToManyDemo() {
	n, itemCount, pageSize := 20, 30000, 1000
	var pageCount int
	if itemCount != 0 {
		if itemCount%pageSize == 0 {
			pageCount = itemCount / pageSize
		} else {
			pageCount = (itemCount / pageSize) + 1
		}
	}

	eg, _ := errgroup.WithContext(context.TODO())
	producerFunc := func() error {
		Producer(pageSize, pageCount)
		return nil
	}

	eg.Go(producerFunc)

	for i := 0; i < n; i++ {
		tmp := i
		eg.Go(func() error {
			return Consumer(tmp)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(errors.WithStack(err))
	} else {
		fmt.Println("success")
	}
}
