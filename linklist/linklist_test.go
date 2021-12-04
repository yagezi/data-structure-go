package linklist

import (
	"fmt"
	"testing"
)

func Test_append(t *testing.T) {
	l := &LinkList{}
	for _, v := range []string{"0", "1", "2", "3", "4", "5", "6"} {
		Append(l, v)
	}

	for p := l.head; p != nil; p = p.next {
		fmt.Printf("%v\n", p.elem)
	}
}

func Test_delete(t *testing.T) {
	l := &LinkList{}
	for _, v := range []string{"0", "1", "2", "3", "4", "5", "6"} {
		Append(l, v)
	}
	elem := DeleteByIndex(l, 0)
	fmt.Println("elem: ", elem)
	for p := l.head; p != nil; p = p.next {
		fmt.Printf("%v\n", p.elem)
	}

}

func Test_revert_recur(t *testing.T) {
	l := &LinkList{}
	for _, v := range []string{"1", "2", "3", "4", "5"} {
		Append(l, v)
	}
	RecurRevert(l)
	for p := l.head; p != nil; p = p.next {
		fmt.Printf("%v\n", p.elem)
	}

}

func Test_revert_loop(t *testing.T) {
	l := &LinkList{}
	for _, v := range []string{"1", "2", "3", "4", "5"} {
		Append(l, v)
	}
	LoopRevert(l)
	for p := l.head; p != nil; p = p.next {
		fmt.Printf("%v\n", p.elem)
	}

}
