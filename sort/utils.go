/*
*@Time        : 2021/12/03
*@File        : utils.go
*@Description : infrastruct for sorting
 */

package sort

import (
	"fmt"
	"math"
)

type List struct {
	r      []int
	length int
}

func (l List) Print() {
	fmt.Printf("%v\n", l.r[1:])
}

func swap(list *List, i, j int) {
	temp := list.r[i]
	list.r[i] = list.r[j]
	list.r[j] = temp
}

func maxium(list *List) int {
	max := -1
	for _, v := range list.r {
		if v > max {
			max = v
		}
	}
	return max
}

type Listf struct {
	r []float32
}

func (l Listf) Print() {
	fmt.Printf("%v\n", l.r[1:])
}

func maxf(list *Listf) float32 {
	var max float32
	for _, v := range list.r {
		if v > max {
			max = v
		}
	}
	return max
}

func minf(list *Listf) float32 {
	min := float32(math.MaxFloat32)
	for _, v := range list.r {
		if v < min {
			min = v
		}
	}
	return min
}
