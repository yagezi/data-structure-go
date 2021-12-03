package sort_user

import "fmt"

func BubbleSort(list *List) {
	for i := 1; i < list.length-1; i++ {
		for j := list.length - 1; j > i; j-- {
			if list.r[j] < list.r[j-1] {
				swap(list, j, j-1)
			}
		}
	}
}

func BubbleSort2(list *List) {
	isSwap := true
	for i := 1; i < list.length-1 && isSwap; i++ {
		for j := list.length - 1; j > i; j-- {
			if list.r[j] < list.r[j-1] {
				swap(list, j, j-1)
				isSwap = true
			}
		}
	}
}

func SelectSort(list *List) {
	for i := 1; i < list.length; i++ {
		min := i
		for j := i + 1; j < list.length; j++ {
			if list.r[j] < list.r[i] {
				min = j
			}
		}
		if min != i {
			swap(list, i, min)
		}
	}
}

func InsertSort(list *List) {
	var i, j int
	for i = 2; i < list.length; i++ {
		if list.r[i] < list.r[i-1] {
			list.r[0] = list.r[i]
			for j = i - 1; list.r[j] > list.r[0]; j-- {
				list.r[j+1] = list.r[j]
			}
			list.r[j+1] = list.r[0]
		}
	}
}

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
	MSort(list, 1, list.length-1)
}

func MSort(list *List, s, n int) {
	if s == n {
		return
	}
	m := (n + s) / 2
	MSort(list, s, m)
	MSort(list, m+1, n)
	Merge(list, s, m, n)
}

func Merge(list *List, s, m, n int) {
	r2 := make([]int, n-s+1)
	i := s
	j := m + 1
	k := 0
	for ; i <= m && j <= n; k++ {
		if list.r[i] > list.r[j] {
			r2[k] = list.r[j]
			j++
		} else {
			r2[k] = list.r[i]
			i++
		}
	}

	if i <= m {
		for ; k < len(r2); k++ {
			r2[k] = list.r[i]
			i++
		}
	}

	if j <= n {
		for ; k < len(r2); k++ {
			r2[k] = list.r[j]
			j++
		}
	}
	fmt.Println("r2: ", r2)
	copy(list.r[s:n+1], r2[:])
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
