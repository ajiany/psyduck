package main

import (
	"testing"
)

func TestNewQueue(t *testing.T) {

	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	for {
		curNode := queue.Pop()
		if curNode == nil {
			t.Log("ending...")
			return
		}
		t.Logf("curNode %v curQueue length %v", curNode.Val, queue.Length)
	}

}
