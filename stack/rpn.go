package stack

import (
	"errors"
	"fmt"
)

type Notation string

func Priority(a rune) int {
	var OPERATORS = map[rune]int{
		'(': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	return OPERATORS[a]
}

func Mid2Post(mid Notation, stack IStack) (post Notation) {
	if len(mid) == 0 {
		return
	}

	// stack := NewStack(len(mid))
	p := make([]rune, 0, len(mid))

	for _, s := range mid {
		fmt.Println(Notation(s))

		switch s {
		default:
			p = append(p, s)
		case ')':
			t, err := stack.Pop()
			for !errors.Is(err, ErrEmptyStack{}) && t.(rune) != '(' {
				p = append(p, t.(rune))
				t, err = stack.Pop()
			}
		case '(':
			stack.Push(s)
		case '+':
			fallthrough
		case '-':
			fallthrough
		case '*':
			fallthrough
		case '/':
			t, err := stack.Top()
			for !errors.Is(err, ErrEmptyStack{}) && Priority(t.(rune)) >= Priority(s) {
				t, _ = stack.Pop()
				p = append(p, t.(rune))
				t, err = stack.Top()
			}
			stack.Push(s)
		}
	}

	for t, err := stack.Pop(); !errors.Is(err, ErrEmptyStack{}); t, err = stack.Pop() {
		p = append(p, t.(rune))
	}
	return Notation(p)
}
