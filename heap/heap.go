/*
*@Time        : 2022/01/11
*@File        : heap.go
*@Description :
 */

package heap

import "fmt"

type Heap struct {
	size     int
	cmp      CmpFunc
	operator Operator
	data     []interface{}
}

func NewHeap(opts ...OptionFunc) *Heap {
	o := NewOption(opts)
	return &Heap{
		size:     0,
		cmp:      o.cmp,
		operator: o.operator,
		data:     make([]interface{}, 1, o.cap),
	}
}

func (h Heap) Size() int {
	return h.size
}

func (h Heap) Top() interface{} {
	if h.size == 0 {
		return nil
	}
	return h.data[1]
}

func (h *Heap) Push(e interface{}) {
	h.data = append(h.data, e)
	h.size++
	i := len(h.data) - 1
	for (i/2) > 0 && h.operator(h.cmp(h.data[i], h.data[i/2])) {
		h.swap(i/2, i)
		i /= 2
	}
}

func (h *Heap) Pop() interface{} {
	if h.size == 0 {
		return nil
	}
	e := h.data[1]
	n := len(h.data)
	h.swap(1, n-1)
	h.data = h.data[:n-1]
	n--
	h.size--
	// adjust
	i := 1
	j := 2 * i
	for j < n {
		if j+1 < n && h.operator(h.cmp(h.data[j+1], h.data[j])) {
			j++
		}
		if h.operator(h.cmp(h.data[j], h.data[i])) {
			h.swap(i, j)
			i = j
			j *= 2
		} else {
			break
		}
	}
	return e
}

func (h Heap) IsEmpty() bool {
	return h.size == 0
}

func (h Heap) String() string {
	s := fmt.Sprintln("size: ", h.size)
	return fmt.Sprintf("%s %v\n", s, h.data)
}

func (h *Heap) swap(i, j int) {
	tmp := h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = tmp
}
