/*
*@Time        : 2021/12/03
*@File        : compare.go
*@Description : compare sorting algorithm
 */

package sort

import (
	"fmt"
)

func HeapSort(list *List) {
	// 先构建大顶堆
	HeapAdjust(list)

	for i := list.length - 1; i > 1; i-- {
		// 交换根节点和末尾节点
		swap(list, 1, i)
		// 将堆中剩余节点调整为大顶堆
		Adjust(list, 1, i-1)
	}

}

func HeapAdjust(list *List) {
	// 从最后一个叶子节点的父节点开始遍历
	for i := list.length / 2; i > 0; i-- {
		Adjust(list, i, list.length-1)
	}
}
func Adjust(list *List, s, m int) {
	// 找到左子节点
	j := 2 * s
	// 如果没有叶子节点
	if j > m {
		return
	}
	// 找到值最大的子节点
	if j < m && list.r[j] < list.r[j+1] {
		j++
	}
	// 如果父节点值小于子节点，交换，并调整子节点
	if list.r[s] < list.r[j] {
		swap(list, s, j)
		Adjust(list, j, m)
	}
}

func MergeSort(list *List) {
	if list.length == 1 {
		return
	}
	extArr := &List{
		r:      make([]int, list.length),
		length: list.length,
	}

	MSort(list, extArr, 1, list.length-1)
}

func MSort(list, extArr *List, s, n int) {
	if s == n {
		return
	}
	m := (n + s) / 2
	MSort(list, extArr, s, m)
	MSort(list, extArr, m+1, n)
	Merge(list, extArr, s, m, n)
}

func Merge(list, extArr *List, s, m, n int) {
	// r2 := make([]int, n-s+1)
	i := s
	j := m + 1
	k := s
	for ; i <= m && j <= n; k++ {
		if list.r[i] > list.r[j] {
			extArr.r[k] = list.r[j]
			j++
		} else {
			extArr.r[k] = list.r[i]
			i++
		}
	}

	if i <= m {
		for ; k <= n; k++ {
			extArr.r[k] = list.r[i]
			i++
		}
	}

	if j <= n {
		for ; k <= n; k++ {
			extArr.r[k] = list.r[j]
			j++
		}
	}
	fmt.Println("external array: ", extArr.r)
	copy(list.r[s:n+1], extArr.r[s:n+1])
}

func QuickSort(list *List) {
	QSort(list, 1, list.length)
}

func QSort(list *List, s, m int) {
	if s == m {
		return
	}
	pivotIndex := Partition(list, s, m)
	QSort(list, s, pivotIndex)
	QSort(list, pivotIndex+1, m)

}

func Partition(list *List, low, high int) int {
	pivot := list.r[low]
	for low < high {
		// 使用 >= 的问题：会将连续出现的值归为其中一方，使子树不平衡
		for low < high && list.r[high] >= pivot {
			high--
		}
		swap(list, low, high)
		for low < high && list.r[low] <= pivot {
			low++
		}
		swap(list, low, high)
	}
	return low
}

func PartitionAlter(list *List, low, high int) int {
	i, j := low, high
	pivot := list.r[i]
	for i < j {
		for i < j && list.r[j] > pivot {
			j--
		}
		for i < j && list.r[i] < pivot {
			i++
		}
		swap(list, i, j)
		i++
		j--
	}
	swap(list, low, j)
	return i
}

func QuickSort3(list *List) {

}

func QSort3(list *List, low, high int) {
	if low >= high {
		return
	}
	lt, gt := Partition3(list, low, high)
	QSort3(list, low, lt-1)
	QSort3(list, gt, high)
}

func Partition3(list *List, low, high int) (int, int) {
	pivot := list.r[low]
	le := low
	gt := high + 1
	i := low
	for i < gt {
		if list.r[i] > pivot {
			swap(list, i, gt-1)
			gt--
		} else if list.r[i] < pivot {
			swap(list, i, le)
			i++
			le++
		} else {
			i++
		}
	}
	return le, gt
}


