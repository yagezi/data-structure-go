package stack

import "fmt"

type Notation string

type Oprator rune

// Compare returns an integer comparing oprator priority.
// The result will be 0 if self==other, -1 if self < other, and +1 if self > other.
func (o Oprator) ComparePriority(other Oprator) int {
	lhs := OPERATORS[o]
	rhs := OPERATORS[other]
	return lhs - rhs
}

var OPERATORS = map[Oprator]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

func Mid2Post(mid Notation) (post Notation) {
	if len(mid) == 0 {
		return
	}

	stack := NewStack(len(mid))
	p := make([]rune, 0, len(mid))

	for _, s := range mid {
		fmt.Println(string(s))

		switch Oprator(s) {
		default:
			p = append(p, s)
		case Oprator(')'):
			for t := stack.Pop().(Oprator); t != '('; {
				p = append(p, rune(t))
			}
		case Oprator('('):
			stack.Push(s)
		case Oprator('+'):
			fallthrough
		case Oprator('-'):
			fallthrough
		case Oprator('*'):
			fallthrough
		case Oprator('/'):
			if stack.GetTop() > -1 {
				for t := stack.Top().(Oprator); t.ComparePriority(Oprator(s)) >= 0; {
					p = append(p, rune(t))
					stack.Pop()
				}
			}
			stack.Push(Oprator(s))
		}

		// fmt.Println(string(p))
		fmt.Println(stack.data)

	}

	for stack.GetTop() != -1 {
		p = append(p, stack.Pop().(rune))
	}
	return Notation(string(p))
}
