/*
*@Time        : 2022/01/11
*@File        : priority_queue.go
*@Description : implements priority queue with heap
 */

package queue

import "github.com/yagezi/data-structure-go/heap"

type PQueue struct {
	h heap.Heap
}

func NewPQueue(t string) *PQueue {
	var opt heap.OptionFunc
	switch t {
	default:
		opt = heap.MinHeap()
	case "max":
		opt = heap.MaxHeap()
	}
	return &PQueue{
		h: *heap.NewHeap(opt),
	}
}

func (pq PQueue) Size() int {
	return pq.h.Size()
}

func (pq PQueue) Head() interface{} {
	return pq.h.Top()
}

func (pq *PQueue) PushBack(e interface{}) {
	pq.h.Push(e)
}

func (pq *PQueue) PopHead() interface{} {
	return pq.h.Pop()
}

func (pq PQueue) IsEmpty() bool {
	return pq.h.IsEmpty()
}

func (pq PQueue) String() string {
	return pq.h.String()
}
