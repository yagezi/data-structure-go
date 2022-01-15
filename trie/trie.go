/*
*@Time        : 2022/01/05
*@Creator     :
*@File        : trie.go
*@Description : dictionary tree
 */

package trie

import (
	"strings"

	"github.com/yagezi/data-structure-go/stack"
)

type Trie struct {
	root     *Node
	keyCount int
}

func NewTrie() *Trie {
	return &Trie{
		root:     NewNode(),
		keyCount: 0,
	}
}

func (t *Trie) Add(key string) {
	node := t.root
	for _, k := range key {
		node.terms++
		if n, ok := node.children[k]; ok {
			node = n
		} else {
			node = node.NewChild(k, "", false)
		}
	}
	if !node.isTerm {
		node.SetTerm(true)
		node.SetPath(key)
		t.keyCount++
	}
}

func (t *Trie) Search(p string) []string {
	node := findNode(t.root, []rune(p))
	if node == nil {
		return nil
	}
	return traverse(node)
}

func (t *Trie) HasKeyWithPrefix(p string) bool {
	if node := findNode(t.root, []rune(p)); node == nil {
		return false
	}
	return true
}

func (t *Trie) String() string {
	if t.root == nil {
		return ""
	}
	return strings.Join(traverse(t.root), ", ")
}

func findNode(node *Node, runes []rune) *Node {
	if node == nil {
		return nil
	}

	if len(runes) == 0 {
		return node
	}

	n, ok := node.children[runes[0]]
	if !ok {
		return nil
	}

	var nrunes []rune
	if len(runes) > 1 {
		nrunes = runes[1:]
	} else {
		// len(runes) == 1
		return n
	}

	return findNode(n, nrunes)
}

// TODO discard node.path
func traverse(node *Node) []string {
	keys := make([]string, 0, node.terms)
	runes := make([]rune, 0, 100)
	var pos int
	s := stack.NewStack()
	for _, c := range node.children {
		s.Push(c)
	}
	for !s.IsEmpty() {
		m := s.Pop()
		n := m.(*Node)
		runes = append(runes, n.val)
		for _, c := range n.children {
			s.Push(c)
		}
		if n.isTerm {
			keys = append(keys, string(runes))
			if n.terms > 0 {
				pos = n.depth
			} else {
				runes = runes[0:pos]
			}
		}
	}
	return keys
}
