package stack

import "errors"

type Queue struct {
	in     *Stack
	out    *Stack
	length int
}

func NewQueue(length_ int) *Queue {
	return &Queue{
		in:     NewStack(length_),
		out:    NewStack(length_),
		length: length_,
	}
}

func (q *Queue) Push(e Element) error {
	errIn := q.in.Push(e)
	_, errOut := q.out.Top()
	if isFullStack(errIn) && isEmptyStack(errOut) {
		if err := q.roll(); err != nil {
			return err
		}
		errIn = q.in.Push(e)
	}
	return errIn
}

func (q *Queue) Pop() (Element, error) {
	res, errOut := q.out.Pop()
	_, errIn := q.in.Top()
	if isEmptyStack(errOut) && !isEmptyStack(errIn) {
		if err := q.roll(); err != nil {
			return nil, err
		}
		res, errOut = q.out.Pop()
	}
	return res, errOut
}

func (q *Queue) roll() error {
	for t, err := q.in.Pop(); !isEmptyStack(err); t, err = q.in.Pop() {
		if err := q.out.Push(t); err != nil {
			return err
		}
	}
	return nil
}

func isEmptyStack(err error) bool {
	return errors.Is(err, ErrEmptyStack{})
}

func isFullStack(err error) bool {
	return errors.Is(err, ErrFullStack{})
}
