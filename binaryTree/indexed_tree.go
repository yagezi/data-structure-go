/*
*@Time        : 2022/01/28
*@File        : index_tree.go
*@Description : binary index tree, a.k.a Fenwick tree
 */

package binarytree

type FTree struct {
	n   int
	bit []int
}

func NewFTree(arr []int) *FTree {
	ft := &FTree{
		n:   len(arr),
		bit: make([]int, len(arr)+1),
	}
	copy(ft.bit[1:], arr)
	for i := 1; i <= ft.n; i++ {
		j := i + lowbit(i)
		if j <= ft.n {
			ft.bit[j] += ft.bit[i]
		}
	}
	return ft

}

func (ft *FTree) add(idx, x int) {
	for i := idx; i <= ft.n; i += lowbit(i) {
		ft.bit[i] += x
	}
}

func (ft *FTree) sum(idx int) int {
	var ans int
	for i := idx; i > 0; i -= lowbit(i) {
		ans += ft.bit[i]
	}
	return ans
}

func lowbit(x int) int {
	return x & -x
}
