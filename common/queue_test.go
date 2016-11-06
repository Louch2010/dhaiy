package common

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 100; i++ {
		queue.Offer(i)
	}
	for !queue.IsEmpty() {
		fmt.Println(queue.Pop())
	}
}
