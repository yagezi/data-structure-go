package ds_string

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKMP(t *testing.T) {
	Convey("kmp", t, func() {
		p := NewKmpPattern("abc")

		Convey("matched", func() {
			idx := p.Match("ababcab")

			So(idx, ShouldEqual, 2)
		})

		Convey("not matched", func() {
			idx := p.Match("ababxabx")

			So(idx, ShouldEqual, -1)
		})

		Convey("matched all", func() {
			idxs := p.MatchAll("ababcdefabca")

			So(idxs, ShouldContain, 2)
			So(idxs, ShouldContain, 8)
		})

		Convey("not matched all", func() {
			idxs := p.MatchAll("ababababababx")

			So(idxs, ShouldBeEmpty)
		})
	})
}

func TestSunday(t *testing.T) {
	sp := NewSundayPattern("ax")
	fmt.Println(sp.Match("acbaxaaaax"))
	fmt.Println(sp.MatchAll("acbaxaaaax"))
}
