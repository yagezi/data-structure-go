package stack

type Element interface{}

type Stack struct {
	data   []Element
	top    int
	length int
}

func NewStack(length_ int) *Stack {
	return &Stack{
		data:   make([]Element, length_),
		top:    -1,
		length: length_,
	}
}

func (s *Stack) Pop() (res Element) {
	if s.top == -1 {
		panic("empty stack")
	}
	res = s.data[s.top]
	s.top--
	return
}

func (s *Stack) Push(e Element) {
	if s.top >= s.length-1 {
		panic("full stack")
	}
	s.top++
	s.data[s.top] = e
}

func (s Stack) Top() Element {
	if s.top == -1 {
		panic("empty stack")
	}
	return s.data[s.top]
}

func (s Stack) GetTop() int {
	return s.top
}
