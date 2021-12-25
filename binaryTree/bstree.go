/*
*@Time        : 2021/12/22
*@Creator     :
*@File        : bstree.go
*@Description : 二叉搜索树 binary search tree
 */

package binarytree

import (
	"fmt"
	"strings"
)

type CmpFunc func(a, b interface{}) int

func CmpString(a, b interface{}) int {
	return strings.Compare(a.(string), b.(string))
}

func CmpInt(a, b interface{}) int {
	return a.(int) - b.(int)
}

type BSTree struct {
	head *BiTNode
	cmp  CmpFunc
}

func NewBSTree(f CmpFunc) *BSTree {
	return &BSTree{
		head: nil,
		cmp:  f,
	}

}

func (bst BSTree) Get(key interface{}) interface{} {
	if isFound, cur, _ := bst.searchBST(key); isFound {
		return cur.value
	}
	return nil
}

func (bst *BSTree) Insert(key, val interface{}) {
	isFound, cur, pre := bst.searchBST(key)
	if isFound {
		cur.value = val
		return
	}
	n := &BiTNode{
		key:   key,
		value: val,
		left:  nil,
		right: nil,
	}
	if pre == nil {
		bst.head = n
	} else if bst.cmp(key, pre.key) > 0 {
		pre.right = n
	} else {
		pre.left = n
	}
}

func (bst *BSTree) Delete(key interface{}) bool {
	isFound, cur, pre := bst.searchBST(key)
	if !isFound {
		return false
	}
	if cur.left == nil {
		if pre == nil {
			// cur == root
			bst.head = cur.right
		} else {
			if pre.left == cur {
				pre.left = cur.right
			} else {
				pre.right = cur.right
			}
		}
	} else if cur.right == nil {
		if pre == nil {
			// cur == root
			bst.head = cur.left
		} else {
			if pre.left == cur {
				pre.left = cur.left
			} else {
				pre.right = cur.left
			}
		}
	} else {
		// 左右子树都不为空。交换右子树的最小值
		temp := new(BiTNode)
		minRight := cur.right
		for minRight.left != nil {
			minRight = minRight.left
		}
		temp.key = minRight.key
		temp.value = minRight.value
		if !bst.Delete(minRight.key) {
			panic("delete minRight error")
		}
		cur.key = temp.key
		cur.value = temp.value
	}
	return true
}

func (bst BSTree) String() string {
	if bst.head == nil {
		return ""
	}
	return bstInOrder(bst.head)
}

// searchBST 返回是否找到key所在节点、节点指针、父节点指针。
// 如果head = nil，返回false, nil, nil
func (bst *BSTree) searchBST(key interface{}) (bool, *BiTNode, *BiTNode) {
	if bst.head == nil {
		return false, nil, nil
	}
	var pre *BiTNode
	cur := bst.head
	for cur != nil {
		if bst.cmp(key, cur.key) > 0 {
			pre = cur
			cur = cur.right
		} else if bst.cmp(key, cur.key) < 0 {
			pre = cur
			cur = cur.left
		} else {
			return true, cur, pre
		}
	}
	return false, nil, pre
}

func bstInOrder(node *BiTNode) (ctx string) {
	if node == nil {
		return
	}
	ctx += bstInOrder(node.left)
	ctx += fmt.Sprintf("%v ", node.key)
	ctx += bstInOrder(node.right)
	return
}
