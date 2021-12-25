package skiplist

import (
	"fmt"
	"testing"
)

func TestInsertAndGet(t *testing.T) {
	sl := NewSkipList(
		WithCompareFunc(CmpString),
		WithMaxLevel(5),
		// WithPrintDetail(),
	)

	count := 2000
	for i := 0; i < count; i++ {
		sl.Insert(fmt.Sprint(i), i)
	}
	fmt.Print(sl)

	fmt.Println(sl.Get("1234"))
}

func TestRemove(t *testing.T) {
	sl := NewSkipList(
		WithCompareFunc(CmpString),
		WithMaxLevel(5),
		// WithPrintDetail(),
	)
	count := 10
	for i := 0; i < count; i++ {
		sl.Insert(fmt.Sprint(i), i)
	}
	fmt.Println(sl.Get("9"))
	fmt.Println(sl.Remove("9"))
	fmt.Println(sl.Get("9"))
	fmt.Println(sl.Remove("9"))
}
