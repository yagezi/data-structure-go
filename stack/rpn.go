package stack

import (
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

	p := make([]rune, 0, len(mid))

	for _, s := range mid {
		fmt.Println(Notation(s))

		switch s {
		default:
			p = append(p, s)
		case ')':
			t := stack.Pop()
			for !stack.IsEmpty() && t.(rune) != '(' {
				p = append(p, t.(rune))
				t = stack.Pop()
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
			t := stack.Top()
			for !stack.IsEmpty() && Priority(t.(rune)) >= Priority(s) {
				t = stack.Pop()
				p = append(p, t.(rune))
				t = stack.Top()
			}
			stack.Push(s)
		}
	}

	for t := stack.Pop(); !stack.IsEmpty(); t = stack.Pop() {
		p = append(p, t.(rune))
	}
	return Notation(p)
}
