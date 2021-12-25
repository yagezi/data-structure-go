package stack

type Element interface{}

type IStack interface {
	Pop() (Element, error)
	Top() (Element, error)
	Push(e Element) error
	IsEmpty() bool
}

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

func (s *Stack) Pop() (Element, error) {
	if s.top == -1 {
		return nil, ErrEmptyStack{}
	}
	res := s.data[s.top]
	s.top--
	return res, nil
}

func (s *Stack) Push(e Element) error {
	if s.top >= s.length-1 {
		return ErrFullStack{}
	}
	s.top++
	s.data[s.top] = e
	return nil
}

func (s Stack) Top() (Element, error) {
	if s.top == -1 {
		return nil, ErrEmptyStack{}
	}
	return s.data[s.top], nil
}

func (s Stack) IsEmpty() bool {
	return s.top == -1
}

type ErrEmptyStack struct {
}

func (e ErrEmptyStack) Error() string {
	return "empty stack"
}

type ErrFullStack struct {
}

func (e ErrFullStack) Error() string {
	return "sfull stack"
}
