/*
*@Time        : 2022/01/04
*@Creator     :
*@File        : rbtree.go
*@Description : red black tree
 */

package rbtree

type RBTree struct {
	root *TNode
	size int
	cmp  CmpFunc
}

func NewRBTree(f CmpFunc) *RBTree {
	return &RBTree{
		root: nil,
		size: 0,
		cmp:  f,
	}
}

func (t *RBTree) Size() int {
	return t.size
}

func (t *RBTree) Find(key interface{}) interface{} {
	if n := t.findFirstNode(key); n != nil {
		return n.value
	}
	return nil
}

func (t *RBTree) findFirstNode(key interface{}) *TNode {
	n := t.root
	for n != nil {
		if t.cmp(key, n.key) == 0 {
			return n
		} else if t.cmp(key, n.key) > 0 {
			n = n.right
		} else {
			n = n.left
		}
	}
	return nil
}

func (t *RBTree) Insert(key, val interface{}) {
	x := t.root
	var y *TNode
	for x != nil {
		y = x
		if t.cmp(key, x.key) > 0 {
			x = x.right
		} else if t.cmp(key, x.key) < 0 {
			x = x.left
		} else {
			// key exist
			x.value = val
			return
		}
	}
	n := &TNode{
		parent: y,
		color:  Red,
		key:    key,
		value:  val,
	}
	t.size++
	// if root == nil
	if y == nil {
		n.color = Black
		t.root = n
		return
	}
	if t.cmp(n.key, y.key) > 0 {
		y.right = n
	} else {
		y.left = n
	}
	t.fixupInsert(n)
}

func (t *RBTree) fixupInsert(n *TNode) {
	var uncle *TNode
	// if n.parent is black node, just insert n and stop fixup
	for n.parent != nil && n.parent.color == Red {
		if n.parent == n.parent.parent.left {
			uncle = n.parent.parent.right
			if uncle != nil && n.color == Red {
				n.parent.parent.color = Red
				n.parent.color = Black
				uncle.color = Black
				n = n.parent.parent
			} else {
				if n == n.parent.right {
					n = n.parent
					t.leftRotate(n)
				}
				n.parent.parent.color = Red
				n.parent.color = Black
				t.rigthRotate(n.parent.parent)
			}
		} else {
			uncle = n.parent.parent.left
			if uncle != nil && n.color == Red {
				n.parent.parent.color = Red
				n.parent.color = Black
				uncle.color = Black
				n = n.parent.parent
			} else {
				if n == n.parent.left {
					n = n.parent
					t.rigthRotate(n)
				}
				n.parent.parent.color = Red
				n.parent.color = Black
				t.leftRotate(n.parent.parent)
			}
		}
	}
	// root must black
	t.root.color = Black
}

func (t *RBTree) leftRotate(n *TNode) {
	r := n.right
	n.right = r.left
	if n.right != nil {
		n.right.parent = n
	}
	r.parent = n.parent
	if n.parent == nil {
		t.root = r
	} else if n == n.parent.left {
		n.parent.left = r
	} else {
		// n == n.parent.right
		n.parent.right = r
	}
	r.left = n
	n.parent = r
}

func (t *RBTree) rigthRotate(n *TNode) {
	l := n.left
	n.left = l.right
	if n.left != nil {
		n.left.parent = n
	}
	l.parent = n.parent
	if n.parent == nil {
		t.root = l
	} else if n == n.parent.left {
		n.parent.left = l
	} else {
		n.parent.right = l
	}
	l.right = n
	n.parent = l
}
