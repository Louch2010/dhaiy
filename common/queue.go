package common

import (
	"fmt"
	"sync"
)

//列队
type Queue struct {
	sync.RWMutex
	q    []interface{}
	size int
}

//创建列表
func NewQueue() *Queue {
	queue := &Queue{
		q:    make([]interface{}, 10),
		size: 0,
	}
	return queue
}

//获取队列大小
func (this *Queue) Size() int {
	this.RLock()
	defer this.RUnlock()
	return this.size
}

//队列是否为空
func (this *Queue) IsEmpty() bool {
	return this.Size() == 0
}

//向队列压入元素
func (this *Queue) Offer(e interface{}) {
	this.Lock()
	defer this.Unlock()
	if this.size == len(this.q) {
		fmt.Println("扩容...")
		//扩容到原来的两倍
		ori := this.q
		this.q = make([]interface{}, len(ori)*2)
		for index, o := range ori {
			this.q[index] = o
		}
	}
	this.q[this.size] = e
	this.size++
}

//从队列弹出元素
func (this *Queue) Pop() interface{} {
	this.Lock()
	defer this.Unlock()
	if this.size > 0 {
		e := this.q[0]
		//所有元素向前移动一位
		for index, t := range this.q {
			if index > 0 {
				this.q[index-1] = t
			}
		}
		this.size--
		return e
	}
	return nil
}

//获取队列头元素
func (this *Queue) Peek() interface{} {
	this.RLock()
	defer this.RUnlock()
	if this.size > 0 {
		return this.q[0]
	}
	return nil
}
