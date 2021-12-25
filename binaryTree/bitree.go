package binarytree

type BiTree struct {
	head *BiTNode
}

// NewBiTree create a binary tree by a full binary tree serial
// '#' present nil. first char in seral must be '#'
func NewBiTree(serial string) *BiTree {
	return &BiTree{
		head: createBiTree(serial, 1),
	}
}

func createBiTree(serial string, idx int) *BiTNode {
	if idx >= len(serial) {
		return nil
	}
	e := serial[idx]
	if e == '#' {
		return nil
	}
	return &BiTNode{
		key:   e,
		left:  createBiTree(serial, 2*idx),
		right: createBiTree(serial, 2*idx+1),
	}
}
