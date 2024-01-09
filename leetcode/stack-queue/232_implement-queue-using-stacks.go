package stack_queue

import "sync"

type Node struct {
	Val  int
	Pre  *Node
	Next *Node
}

type MyQueue struct {
	mux    sync.RWMutex
	Length int
	Head   *Node
	Tail   *Node
}

func Constructor() MyQueue {
	return MyQueue{Head: &Node{Val: -1}, Tail: &Node{Val: -1}}
}

func (this *MyQueue) Push(x int) {
	this.mux.Lock()
	defer this.mux.Unlock()
	newNode := Node{Val: x}
	if this.Length == 0 {
		this.Head.Next = &newNode
		newNode.Pre = this.Head.Next
		newNode.Next = this.Tail
		this.Tail.Pre = &newNode
	}

	tailPreNode := this.Tail.Pre
	tailPreNode.Next = &newNode
	newNode.Pre = tailPreNode
	newNode.Next = this.Tail
	this.Tail.Pre = &newNode

	this.Length++
}

func (this *MyQueue) Pop() int {
	this.mux.Lock()
	defer this.mux.Unlock()
	if this.Length == 0 {
		return -1
	}

	headNextNode := this.Head.Next
	this.Head.Next = headNextNode.Next
	headNextNode.Next.Pre = this.Head

	this.Length--
	return headNextNode.Val
}

func (this *MyQueue) Peek() int {
	this.mux.RLock()
	defer this.mux.RUnlock()
	if this.Length == 0 {
		return -1
	}

	return this.Head.Next.Val
}

func (this *MyQueue) Empty() bool {
	this.mux.RLock()
	defer this.mux.RUnlock()

	return this.Length == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
