package stack

import (
	"fmt"
	"testing"
)

func Test_Mid2Post(t *testing.T) {
	mid := Notation("9+(3-1)*3+8/2")
	stack := NewStack()
	fmt.Println(Mid2Post(mid, stack))
}
