package queue

import (
	"container/list"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueue(t *testing.T) {
	Convey("Given a emtpy queue", t, func() {
		q := NewQueue()

		Convey("Then isEmpty is true", func() {
			So(q.IsEmpty(), ShouldBeTrue)
		})

		Convey("Then Length is 0", func() {
			So(q.Length(), ShouldEqual, 0)
		})

		Convey("Then Head is nil", func() {
			So(q.Head(), ShouldBeNil)
		})

		Convey("when add element", func() {
			q.Enqueue(1)

			Convey("Then isEmpty is false", func() {
				So(q.IsEmpty(), ShouldBeFalse)
			})

			Convey("Then Length is 1", func() {
				So(q.Length(), ShouldEqual, 1)
			})

			Convey("Then Head is 1", func() {
				So(q.Head(), ShouldEqual, 1)
			})
		})

		Convey("when add element, and then pop element", func() {
			q.Enqueue(1)
			q.Dequeue()

			Convey("Then isEmpty is true", func() {
				So(q.IsEmpty(), ShouldBeTrue)
			})

			Convey("Then Length is 0", func() {
				So(q.Length(), ShouldEqual, 0)
			})

			Convey("Then Head is nil", func() {
				So(q.Head(), ShouldBeNil)
			})
		})

		Convey("when add element, and then pop element twice", func() {
			q.Enqueue(1)
			q.Dequeue()
			q.Dequeue()

			Convey("Then isEmpty is true", func() {
				So(q.IsEmpty(), ShouldBeTrue)
			})

			Convey("Then Length is 0", func() {
				So(q.Length(), ShouldEqual, 0)
			})

			Convey("Then Head is nil", func() {
				So(q.Head(), ShouldBeNil)
			})
		})
	})

}

func Test_PriorityQueue(t *testing.T) {
	pq := NewPQueue("min")
	for _, v := range []int{5, 3, 4, 1, 2} {
		pq.PushBack(v)
	}
	fmt.Println(pq)
	fmt.Println("head is ", pq.Head())

	for !pq.IsEmpty() {
		fmt.Println(pq.PopHead())
	}
}

func Test_MonoQueue(t *testing.T) {
	// 单调队列模板，递增
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	l := list.New()

	for _, v := range arr {
		for l.Len() > 0 && l.Back().Value.(int) >= v {
			l.Remove(l.Back())
		}
		l.PushBack(v)
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v ", e.Value.(int))
		}
		fmt.Println()
	}
}

func Test_MonoQueueIndex(t *testing.T) {
	// 单调队列模板, 索引表示，递增
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// 滑动窗长度
	win := 3
	l := list.New()

	for i := range arr {
		for l.Len() > 0 && arr[l.Back().Value.(int)] <= arr[i] {
			l.Remove(l.Back())
		}
		l.PushBack(i)
		for l.Front().Value.(int) < i-win+1 {
			l.Remove(l.Front())
		}
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v ", arr[e.Value.(int)])
		}
		fmt.Println()
	}

}
