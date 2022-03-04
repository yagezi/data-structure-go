/*
*@Time        : 2022/01/28
*@File        : segment_tree.go
*@Description : segment tree
 */

package binarytree

type STree struct {
	n   int
	sum []int
	add []int
}

func NewSTree(arr []int) *STree {
	n := len(arr)
	st := &STree{
		n:   n,
		sum: make([]int, n<<4),
		add: make([]int, n<<4),
	}
	st.build(arr, 0, len(arr)-1, 1)
	return st
}

// 合并左右子树的区间
func (st *STree) updateParent(idx int) {
	st.sum[idx] = st.sum[idx<<1] + st.sum[idx<<1|1]
}

// build recursively build sigment tree,
// l <= i <= r, left-closed and right-closed interval,
// 2*idx = idx <<1, while 2*idx+1 = idx<<1|1
func (st *STree) build(arr []int, l, r, idx int) {
	if l == r {
		st.sum[idx] = arr[l]
		return
	}
	mid := int(uint(l+r) >> 1)
	st.build(arr, l, mid, idx<<1)
	st.build(arr, mid+1, r, idx<<1|1)
	st.updateParent(idx)
}

// Query implements interval sum query recursively
// query sum(arr[L..R])
func (st *STree) Query(L, R int) int {
	return st.queryR(L, R, 0, st.n-1, 1)
}

// L, R为用户查询区间
// l, r为节点代表的区间
func (st *STree) queryR(L, R, l, r, idx int) int {
	// 如果节点区间在查询区间内，返回节点的值
	if L <= l && r <= R {
		return st.sum[idx]
	}
	mid := int(uint(l+r) >> 1)
	var sum int
	// 否则，遍历左右子树并求和
	if L <= mid {
		sum += st.queryR(L, R, l, mid, idx<<1)
	}

	if R > mid {
		sum += st.queryR(L, R, mid+1, r, idx<<1|1)
	}
	return sum
}

// Update implements point update
func (st *STree) Update(i, val int) {
	st.updateR(i, val, 0, st.n-1, 1)
}

func (st *STree) updateR(i, val, l, r, idx int) {
	if l == r {
		st.sum[idx] = val
		return
	}
	mid := int(uint(l+r) >> 1)
	if i <= mid {
		st.updateR(i, val, l, mid, idx<<1)
	} else {
		st.updateR(i, val, mid+1, r, idx<<1|1)
	}
	// 左右子树有更新，因此也需要更新父节点
	st.updateParent(idx)
}

// TODO
// // AddMulti implements interval add by 'Lazy Propagation'
// // arr[L,R] += val
// func (st *STree) AddMulti(L, R, val int) {

// }

// func (st *STree) addMultiR(L, R, val, l, r, idx int) {
// 	if L <= l && r <= R {
// 		//更新区间和，注意区间内每个元素都增加val
// 		st.sum[idx] += val * (r - l + 1)
// 		//增加Add标记，表示本区间的Sum正确，子区间的Sum仍需要根据Add的值来调整
// 		st.add[idx] += val
// 		return
// 	}
// }
