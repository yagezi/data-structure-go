/*
*@Time        : 2022/01/11
*@File        : heap_test.go
*@Description :
 */
package heap

import (
	"fmt"
	"testing"
)

func Test_heap(t *testing.T) {
	h := NewHeap()
	for _, v := range []int{5, 4, 3, 2, 1} {
		h.Push(v)
	}
	fmt.Println(h)

	h.Pop()
	fmt.Println(h)

	fmt.Println(h.Top())

}

func Test_heapString(t *testing.T) {
	h := NewHeap(MaxHeap(), WithCmpFunc(CmpString))
	for _, v := range []string{"a", "b", "c", "d", "e"} {
		h.Push(v)
	}
	fmt.Println(h)
	fmt.Println(h.Top())

	for _, v := range []string{"a", "b", "c", "d", "e"} {
		h.Push(v)
	}
	h.Pop()
	fmt.Println(h)
}
