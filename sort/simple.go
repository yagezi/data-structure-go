/*
*@Time        : 2021/12/03
*@File        : simple.go
*@Description : simple sorting algorithm
 */

package sort

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
