package stack

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoStackQueue(t *testing.T) {
	Convey("given an empty queue", t, func() {
		queue := NewQueue()

		Convey("when pop", func() {
			e := queue.Pop()

			Convey("then should be empty", func() {
				So(e, ShouldBeNil)
			})
		})
	})
}

func TestQueue(t *testing.T) {
	queue := NewQueue()
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
