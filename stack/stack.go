package stack

type Element interface{}

type IStack interface {
	Pop() Element
	Top() Element
	Push(e Element)
	IsEmpty() bool
}

type Stack struct {
	data []Element
	size int
}

func NewStack(cap ...int) *Stack {
	c := 1
	if len(cap) > 0 {
		c = cap[0]
	}
	return &Stack{
		data: make([]Element, 0, c),
	}
}

func (s *Stack) Pop() Element {
	if s.size == 0 {
		return nil
	}
	res := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return res
}

func (s *Stack) Push(e Element) {
	s.data = append(s.data, e)
	s.size++
}

func (s Stack) Top() Element {
	if s.size == 0 {
		return nil
	}
	return s.data[s.size-1]
}

func (s Stack) IsEmpty() bool {
	return s.size == 0
}
