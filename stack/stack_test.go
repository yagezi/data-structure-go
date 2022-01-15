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
