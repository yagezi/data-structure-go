package stack

import (
	"fmt"
	"testing"
)

func Test_stack(t *testing.T) {
	s := NewStack()
	s.Push("0")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push("1")
	fmt.Println(s.Pop())
}

func Test_MonoStack(t *testing.T) {
	// 单调栈模板，递增
	arr := []int{3, 2, 3, 4, 2, 6, 4, 5, 2, 3}
	s := NewStack()
	for _, v := range arr {
		for !s.IsEmpty() && s.Top().(int) <= v {
			s.Pop()
		}
		s.Push(v)
		fmt.Println(s)
	}
}

func Test_MonoStackIndex(t *testing.T) {
	// 单调栈-索引表示，递增
	arr := []int{3, 2, 3, 4, 2, 6, 4, 5, 2, 3}
	s := NewStack()
	for i := range arr {
		for !s.IsEmpty() && arr[s.Top().(int)] <= arr[i] {
			s.Pop()
		}
		s.Push(i)
		fmt.Println(s)
	}
}
