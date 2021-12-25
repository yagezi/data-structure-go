package queue

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQStack(t *testing.T) {
	Convey("given a new qstack", t, func() {
		qstack := NewQStack()

		Convey("Then isEmpty should be true", func() {
			So(qstack.IsEmpty(), ShouldBeTrue)
		})

		Convey("Length should be 0", func() {
			So(qstack.Length(), ShouldEqual, 0)
		})

		Convey("Top should be nil", func() {
			So(qstack.Top(), ShouldBeNil)
		})

		Convey("When push element twice", func() {
			qstack.Push(1)
			qstack.Push(2)

			Convey("Then isEmpty should be false", func() {
				So(qstack.IsEmpty(), ShouldBeFalse)
			})

			Convey("Length should be 2", func() {
				So(qstack.Length(), ShouldEqual, 2)
			})

			Convey("Top should correct", func() {
				So(qstack.Top(), ShouldEqual, 2)
			})
		})

		Convey("When push element twice, and them", func() {
			qstack.Push(1)
			qstack.Push(2)

			Convey("Then element seqence should correct", func() {
				So(qstack.Pop(), ShouldEqual, 2)
				So(qstack.Pop(), ShouldEqual, 1)

				Convey("isEmpty should be true", func() {
					So(qstack.IsEmpty(), ShouldBeTrue)
				})

				Convey("Length should be 0", func() {
					So(qstack.Length(), ShouldEqual, 0)
				})

				Convey("Top should be nil", func() {
					So(qstack.Top(), ShouldBeNil)
				})
			})
		})
	})
}
