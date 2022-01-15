/*
*@Time        : 2021/12/03
*@File        : sort_test.go
*@Description : sorting unittest
 */

package sort

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
	extArr := &List{
		r:      make([]int, list.length),
		length: list.length,
	}
	Merge(list, extArr, 1, list.length/2, list.length-1)
	list.Print()
}

func Test_MSort(t *testing.T) {
	list := &List{
		r:      []int{-1, 4, 3, 2, 1},
		length: 5,
	}
	extArr := &List{
		r:      make([]int, list.length),
		length: list.length,
	}
	MSort(list, extArr, 1, 4)
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
		r:      []int{-1, 5, 1, 2, 3, 4, 5, 5, 5, 8},
		length: 9,
	}
	Partition(list, 1, 9)
	list.Print()
}

func Test_PartitionAlter(t *testing.T) {
	list := &List{
		r:      []int{-1, 5, 5, 2, 3, 4, 5, 2, 5, 8},
		length: 9,
	}
	PartitionAlter(list, 1, 9)
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

func Test2_Partition3(t *testing.T) {
	list := &List{
		r:      []int{-1, 3, 4, 2},
		length: 4,
	}
	fmt.Println(Partition3(list, 1, 3))
	list.Print()
}

func Test_CountSort(t *testing.T) {
	list := &List{
		r:      []int{9, 3, 5, 4, 9, 1, 2, 7, 8, 1, 3, 6, 5, 3, 4, 0, 10, 9, 7, 9},
		length: 20,
	}
	CountSort(list)
	list.Print()
}

func Test_RadixSort(t *testing.T) {
	list := &List{
		r:      []int{170, 45, 75, 90, 802, 24, 2, 66},
		length: 8,
	}
	RadixSort(list)
	list.Print()
}

func Test_BucketSort(t *testing.T) {
	list := &Listf{
		r: []float32{9.8, 0.6, 10.1, 1.9, 3.07, 3.04, 5.0, 8.0, 4.8, 7.68},
	}
	BucketSort(list, 5)
	list.Print()
}

func Test_floatInserttSort(t *testing.T) {
	list := &Listf{
		r: []float32{9.8, 0.6, 10.1, 1.9, 3.07, 3.04, 5.0, 8.0, 4.8, 7.68},
	}
	fmt.Println(insertSortFoalt(list.r))
}
