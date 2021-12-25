package queue

import (
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
