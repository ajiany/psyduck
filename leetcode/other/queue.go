package main

import "sync"

func main() {
	mmap := make(map[string]interface{}, 0)

	mux := sync.RWMutex{}
	readFunc := func(idx string) interface{} {
		mux.RLock()
		defer mux.RUnlock()
		return mmap[idx]
	}

	writeFunc := func(idx string, val interface{}) {
		mux.Lock()
		defer mux.Unlock()
		mmap[idx] = val
	}

	readFunc("1")
	writeFunc("1", 1)

}

// Queue 队列
type Queue struct {
	TopNode  *Node
	TailNode *Node
	Length   int
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.TopNode = &Node{
		Val: -1,
	}
	queue.TailNode = &Node{
		Val: -1,
	}
	queue.Length = 0

	return queue
}

type Node struct {
	Val      int
	PreNode  *Node
	NextNode *Node
}

func (queue *Queue) Leng() int {

	return queue.Length
}

func (queue *Queue) Pop() *Node {

	// 队列为空则提前返回
	if queue.Length == 0 {
		return nil
	}

	// 当前队列的头节点
	oldPreNode := queue.TopNode.NextNode
	queue.TopNode.NextNode = oldPreNode.NextNode
	if oldPreNode.NextNode != nil {
		oldPreNode.NextNode.PreNode = queue.TopNode
	}

	queue.Length = queue.Length - 1

	return oldPreNode
}

func (queue *Queue) Push(val int) {
	curNode := &Node{
		Val: val,
	}
	// 判断当前队列是否为空
	if queue.Length == 0 {

		// 当前节点的前置节点为头节点
		curNode.PreNode = queue.TopNode
		curNode.NextNode = queue.TailNode

		queue.TopNode.NextNode = curNode
		queue.TailNode.PreNode = curNode
	} else {
		// 队列不为空
		oldPreNode := queue.TailNode.PreNode
		oldPreNode.NextNode = curNode
		curNode.PreNode = oldPreNode

		// 先将尾节点指向新节点
		queue.TailNode.PreNode = curNode
		curNode.NextNode = queue.TailNode

	}

	queue.Length = queue.Length + 1

	return
}
