package binarytree

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinaryTree(t *testing.T) {
	tree := NewBiTree("#eag#cf###bd")
	PrevOrderRecurse(tree)
	PreOrderStack(tree)

	InOrderRecurse(tree)
	InOrderStack(tree)

	PostOrderRecurse(tree)

	LevelOrderQueue(tree)
}

func TestBinarySearchTree(t *testing.T) {
	Convey("bst", t, func() {
		bst := NewBSTree(CmpInt)

		Convey("empty get", func() {
			So(bst.Get("empty"), ShouldBeNil)
		})

		Convey("empty delete", func() {
			So(bst.Delete("empty"), ShouldBeFalse)
		})

		Convey("insert insert", func() {
			bst.Insert(1, 1)
			So(bst.String(), ShouldEqual, "1 ")
		})

		Convey("get head", func() {
			bst.Insert(1, 1)
			So(bst.Get(1), ShouldEqual, 1)
		})

		Convey("delete head", func() {
			bst.Insert(1, 1)
			bst.Delete(1)
			So(bst.String(), ShouldBeEmpty)
		})

		Convey("basic delete", func() {
			bst.Insert(1, 1)
			bst.Insert(2, 2)
			bst.Delete(2)
			So(bst.Get(2), ShouldBeNil)
		})

		Convey("basic get", func() {
			bst.Insert(1, 1)
			bst.Insert(2, 2)
			So(bst.Get(2), ShouldEqual, 2)
		})

		Convey("insert existed head key", func() {
			bst.Insert(1, 1)
			bst.Insert(1, 10)
			So(bst.Get(1), ShouldEqual, 10)
		})

		Convey("insert existed key", func() {
			bst.Insert(2, 2)
			bst.Insert(2, 20)
			So(bst.Get(2), ShouldEqual, 20)
		})

		Convey("delete leaf", func() {
			for _, v := range []int{4, 3, 5} {
				bst.Insert(v, v)
			}
			bst.Delete(5)
			So(bst.String(), ShouldEqual, "3 4 ")
			So(bst.Get(5), ShouldBeNil)
		})

		Convey("delete node with empty right", func() {
			for _, v := range []int{62, 58, 47, 88} {
				bst.Insert(v, v)
			}
			bst.Delete(58)
			So(bst.String(), ShouldEqual, "47 62 88 ")
			So(bst.Get(58), ShouldBeNil)
		})

		Convey("delete node with empty left", func() {
			for _, v := range []int{62, 58, 61, 88} {
				bst.Insert(v, v)
			}
			bst.Delete(58)
			So(bst.String(), ShouldEqual, "61 62 88 ")
			So(bst.Get(58), ShouldBeNil)
		})

		Convey("delete node with left and right subtree", func() {
			for _, v := range []int{62, 88, 58, 47, 35, 51, 49, 50} {
				bst.Insert(v, v)
			}
			bst.Delete(58)
			So(bst.String(), ShouldEqual, "35 47 49 50 51 62 88 ")
			So(bst.Get(58), ShouldBeNil)
		})

	})
}

func Test_SegmentTree(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	st := NewSTree(arr)
	fmt.Println(st)

	st.Update(1, 2)
	fmt.Println(st)

	// fmt.Println(st.Query(1, 4))

}

func Test_IndexedTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	ftree := NewFTree(arr)
	fmt.Println(ftree)
	fmt.Println(ftree.sum(4))
	ftree.add(1, 1)
	fmt.Println(ftree)
	fmt.Println(ftree.sum(4))
}
