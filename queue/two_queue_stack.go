package queue

type QStack struct {
	master,
	minion *Queue
}

func NewQStack() *QStack {
	return &QStack{
		master: &Queue{},
		minion: &Queue{},
	}
}

func (qs *QStack) Push(e Element) {
	qs.master.Enqueue(e)
}

func (qs *QStack) Pop() Element {
	if qs.master.IsEmpty() {
		return nil
	}
	return qs.swapAndGet()
}

func (qs QStack) Top() Element {
	return qs.master.Tail()
}

func (qs QStack) IsEmpty() bool {
	return qs.master.IsEmpty()
}

func (qs QStack) Length() int {
	return qs.master.Length()
}

func (qs *QStack) swapAndGet() (res Element) {
	for qs.master.Length() > 1 {
		qs.minion.Enqueue(qs.master.Dequeue())
	}
	res = qs.master.Dequeue()

	temp := qs.master
	qs.master = qs.minion
	qs.minion = temp
	return
}
