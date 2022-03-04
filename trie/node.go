/*
*@Time        : 2022/01/06
*@Creator     :
*@File        : node.go
*@Description : trie / prefix tree / radix tree node
 */

package trie

type Node struct {
	val      rune
	path     string
	depth    int
	isTerm   bool
	terms    int
	children map[rune]*Node
}

func NewNode() *Node {
	n := &Node{
		val:      0,
		path:     "",
		depth:    0,
		isTerm:   false,
		terms:    0,
		children: make(map[rune]*Node),
	}
	return n
}

func (parent *Node) NewChild(val rune, path string, isTerminal bool) *Node {
	n := &Node{
		val:      val,
		path:     path,
		depth:    parent.depth + 1,
		isTerm:   isTerminal,
		terms:    0,
		children: make(map[rune]*Node),
	}
	parent.children[n.val] = n
	return n
}

func (n *Node) SetTerm(isTerm bool) {
	n.isTerm = isTerm
}

func (n *Node) SetPath(key string) {
	n.path = key
}
