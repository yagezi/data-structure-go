package queue

import "fmt"

type Element interface {
}

type Node struct {
	elem Element
	next *Node
}

type IQueue interface {
	Dequeue() Element
	Enqueue(e Element)
	Head() Element
	IsEmpty() bool
	Length() int
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func NewQueue() *Queue {
	return &Queue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *Queue) Dequeue() Element {
	if q.head == nil {
		return nil
	}
	head := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	head.next = nil
	q.length--
	return head.elem
}

func (q *Queue) Enqueue(e Element) {
	if q.head == nil {
		q.head = &Node{
			elem: e,
			next: nil,
		}
		q.tail = q.head
		q.length++
		return
	}
	q.tail.next = &Node{
		elem: e,
		next: nil,
	}
	q.tail = q.tail.next
	q.length++
}

func (q Queue) Head() Element {
	if q.head == nil {
		return nil
	}
	return q.head.elem
}

func (q Queue) Tail() Element {
	if q.tail == nil {
		return nil
	}
	return q.tail.elem
}

func (q Queue) IsEmpty() bool {
	return q.head == nil
}

func (q Queue) Length() int {
	return q.length
}

func (q Queue) print() {
	fmt.Println("length", q.length)
	for p := q.head; p != nil; p = p.next {
		fmt.Printf("%v ", p.elem)
	}
	fmt.Println("")
}
