package stack

import (
	"fmt"
	"testing"
)

func Test_stack(t *testing.T) {
	s := NewStack(1)
	s.Push("0")
	fmt.Println(s.Pop())
}

func Test_Mid2Post(t *testing.T) {
	mid := Notation("9+(3-1)*3+10/2")
	fmt.Println(Mid2Post(mid))
}
