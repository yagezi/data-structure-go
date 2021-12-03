package sort_user

import "fmt"

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
