package stack

type Queue struct {
	in  *Stack
	out *Stack
}

func NewQueue() *Queue {
	return &Queue{
		in:  NewStack(),
		out: NewStack(),
	}
}

func (q *Queue) Push(e Element) {
	q.in.Push(e)
}

func (q *Queue) Pop() Element {
	if q.out.IsEmpty() {
		q.roll()
	}
	return q.out.Pop()
}

func (q *Queue) roll() {
	for !q.in.IsEmpty() {
		q.out.Push(q.in.Pop())
	}
}
