package sort_user

import (
	"fmt"
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	list := &List{
		r:      []int{-1, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		length: 10,
	}
	BubbleSort(list)
	list.Print()
}

func Test_BubbleSort2(t *testing.T) {
	list := &List{
		r:      []int{-1, 2, 1, 3, 4, 5, 6, 7, 8, 9},
		length: 10,
	}
	BubbleSort2(list)
	list.Print()
}

func Test_SelectSort(t *testing.T) {
	list := &List{
		r:      []int{-1, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		length: 10,
	}
	SelectSort(list)
	list.Print()
}

func Test_InsertSort(t *testing.T) {
	list := &List{
		r:      []int{-1, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		length: 10,
	}
	InsertSort(list)
	list.Print()
}

func Test_HeapAdjust(t *testing.T) {
	list := &List{
		r:      []int{-1, 1, 2, 7, 3, 5, 9, 8, 6, 4},
		length: 10,
	}
	HeapAdjust(list)
	list.Print()
}

func Test_HeapSort(t *testing.T) {
	list := &List{
		r:      []int{-1, 1, 2, 7, 3, 5, 9, 8, 6, 4},
		length: 10,
	}
	HeapSort(list)
	list.Print()
}

func Test_Merge(t *testing.T) {
	list := &List{
		r:      []int{-1, 3, 4, 1, 2},
		length: 5,
	}
	Merge(list, 1, list.length/2, list.length-1)
	list.Print()
}

func Test_MSort(t *testing.T) {
	list := &List{
		r:      []int{-1, 4, 3, 2, 1},
		length: 5,
	}
	MSort(list, 1, 4)
	list.Print()
}

func Test_MergeSort(t *testing.T) {
	list := &List{
		r:      []int{-1, 4, 3, 2, 1, 9, 8, 7},
		length: 8,
	}
	MergeSort(list)
	list.Print()
}

func Test_Partition(t *testing.T) {
	list := &List{
		r:      []int{-1, 5, 1, 9, 3, 7, 4, 8, 6, 2},
		length: 9,
	}
	Partition(list, 1, 9)
	list.Print()
}

func Test_QSort(t *testing.T) {
	list := &List{
		r:      []int{-1, 5, 1, 9, 3, 7, 4, 8, 6, 2},
		length: 9,
	}
	QSort(list, 1, 9)
	list.Print()
}

func Test_QuickSort(t *testing.T) {
	list := &List{
		r:      []int{-1, 5, 1, 9, 3, 7, 4, 8, 6, 2},
		length: 9,
	}
	QuickSort(list)
	list.Print()
}

func Test_Partition3(t *testing.T) {
	list := &List{
		r:      []int{-1, 5, 1, 5, 9, 3, 7, 4, 8, 5, 2},
		length: 10,
	}
	fmt.Println(Partition3(list, 1, 10))
	list.Print()
}

func Test_QSort3(t *testing.T) {
	list := &List{
		r:      []int{-1, 5, 1, 5, 9, 3, 7, 4, 8, 5, 2},
		length: 10,
	}
	QSort3(list, 1, 10)
	list.Print()
}

func Test_Partition3Test(t *testing.T) {
	list := &List{
		r:      []int{-1, 3, 4, 2},
		length: 4,
	}
	fmt.Println(Partition3(list, 1, 3))
	list.Print()
}
