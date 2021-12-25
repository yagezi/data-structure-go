package stack

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoStackQueue(t *testing.T) {
	Convey("given an empty queue", t, func() {
		queue := NewQueue(1)

		Convey("when pop", func() {
			_, err := queue.Pop()

			Convey("then should be empty error", func() {
				So(err, ShouldHaveSameTypeAs, ErrEmptyStack{})
			})
		})
	})

	Convey("given an full queue", t, func() {
		queue := NewQueue(2)
		queue.Push("1")
		queue.roll()
		queue.Push("2")
		queue.Push("3")

		Convey("when push", func() {
			err := queue.Push("4")

			Convey("then should be error", func() {
				So(err, ShouldHaveSameTypeAs, ErrFullStack{})
			})
		})
	})

	Convey("given an queue with length of 2", t, func() {
		queue := NewQueue(2)

		Convey("when push 3 times", func() {
			queue.Push("1")
			queue.Push("2")
			queue.Push("3")

			Convey("then should roll, and out stack should be full ", func() {
				So(queue.out.Push("4"), ShouldHaveSameTypeAs, ErrFullStack{})
			})
		})
	})
}

func TestQueue(t *testing.T) {
	queue := NewQueue(2)
	queue.Push("1")
	queue.Push("2")
	queue.Push("3")
	queue.Push("4")
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	queue.Push("1")
	queue.Push("2")
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
