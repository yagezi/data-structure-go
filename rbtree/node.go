/*
*@Time        : 2022/01/04
*@Creator     :
*@File        : node.go
*@Description : red black tree
 */

package rbtree

type Color bool

const (
	Red   = true
	Black = false
)

type TNode struct {
	parent *TNode
	left    *TNode
	right   *TNode
	color   Color
	key     interface{}
	value   interface{}
}

func (tn *TNode) setValue(val interface{}) {
	tn.value = val
}

func (tn *TNode) max() *TNode {
	for tn.right != nil {
		tn = tn.right
	}
	return tn
}

func (tn *TNode) min() *TNode {
	for tn.left != nil {
		tn = tn.left
	}
	return tn
}
