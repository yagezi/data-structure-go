/*
*@Time        : 2022/01/10
*@File        : count.go
*@Description : count sorting algorithm
 */

package sort

func CountSort(list *List) {
	max := maxium(list)
	counts := make([]int, max+1)
	tmp := make([]int, list.length)

	for i := 0; i < list.length; i++ {
		counts[list.r[i]]++
	}

	for i := 1; i <= max; i++ {
		counts[i] += counts[i-1]
	}

	for i := list.length - 1; i >= 0; i-- {
		tmp[counts[list.r[i]]-1] = list.r[i]
		counts[list.r[i]]--
	}

	copy(list.r, tmp)
}

func RadixSort(list *List) {
	max := maxium(list)
	for exp := 1; max/exp > 0; exp *= 10 {
		IntCountSort(list, exp)
	}
}

func IntCountSort(list *List, exp int) {
	max := maxium(list)
	counts := make([]int, max+1)
	tmp := make([]int, list.length)

	for i := 0; i < list.length; i++ {
		counts[list.r[i]/exp%10]++
	}

	for i := 1; i <= max; i++ {
		counts[i] += counts[i-1]
	}

	for i := list.length - 1; i >= 0; i-- {
		tmp[counts[list.r[i]/exp%10]-1] = list.r[i]
		counts[list.r[i]/exp%10]--
	}

	copy(list.r, tmp)
}

func BucketSort(list *Listf, n int) {
	max := maxf(list)
	min := minf(list)
	bucketRange := int(max-min+float32(n)) / n
	// 为方便，此处用slice替代数组+链表
	buckets := make([]([]float32), n)
	for _, v := range list.r {
		idx := int(v-min) / bucketRange
		buckets[idx] = append(buckets[idx], v)
	}
	for i, list := range buckets {
		buckets[i] = insertSortFoalt(list)
	}

	index := 0
	for i := 0; i < n; i++ {
		copy(list.r[index:], buckets[i])
		index += len(buckets[i])
	}
}

func insertSortFoalt(list []float32) []float32 {
	var i, j int
	var tmp float32
	for i = 1; i < len(list); i++ {
		tmp = list[i]
		for j = i - 1; list[j] > tmp && j >= 0; j-- {
			list[j+1] = list[j]
		}
		list[j+1] = tmp
	}
	return list
}
