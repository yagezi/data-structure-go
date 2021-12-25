package stack

import (
	"fmt"
	"testing"
)

func Test_stack(t *testing.T) {
	s := NewStack(1)
	s.Push("0")
	_, err := s.Pop()
	fmt.Println(err)
	_, err = s.Pop()
	fmt.Println(err)
}

type Opt byte

func Test_Generic(t *testing.T) {
	s := NewStack(1)
	s.Push(Opt('1'))
	e, _ := s.Pop()
	fmt.Println(e.(Opt))
}
