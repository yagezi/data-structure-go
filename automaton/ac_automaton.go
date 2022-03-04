/*
*@Time        : 2022/02/12
*@Description : ac automaton, for muti-pattern match
 */
package acautomaton

import (
	"container/list"
	"fmt"
)

type NodeAC struct {
	id       int
	val      byte
	len      int
	isTerm   bool
	term     string
	failed   *NodeAC
	children map[byte]*NodeAC
}

type Automaton struct {
	root  *NodeAC
	count int
}

func NewAutomaton() *Automaton {
	// root.failed point to fake root
	return &Automaton{
		count: 0,
		root: &NodeAC{
			id:  0,
			len: 0,
			failed: &NodeAC{
				id:       -1,
				children: make(map[byte]*NodeAC),
			},
			children: make(map[byte]*NodeAC),
		},
	}
}

func (ac *Automaton) Add(ptn string) {
	node := ac.root
	for i := range ptn {
		n, ok := node.children[ptn[i]]
		if !ok {
			ac.count++
			n = &NodeAC{
				id:       ac.count,
				val:      ptn[i],
				len:      node.len + 1,
				children: map[byte]*NodeAC{},
			}
			node.children[ptn[i]] = n
		}
		node = n
	}
	node.isTerm = true
	node.term = ptn
}

func (ac *Automaton) GenFailurePointer() {
	var fp *NodeAC
	l := list.New()
	l.PushBack(ac.root)
	for l.Len() > 0 {
		n := l.Front().Value.(*NodeAC)
		l.Remove(l.Front())
		for k, c := range n.children {
			fp = n.failed
			fc, ok := fp.children[k]
			for fp != nil && !ok {
				fc, ok = fp.children[k]
				fp = fp.failed
			}
			if fp == nil {
				fc = ac.root
			}
			c.failed = fc
			l.PushBack(c)
		}
	}
}

func (ac *Automaton) Match(tgt string) {
	p := ac.root
	for i := range tgt {
		c, ok := p.children[tgt[i]]
		for p != nil && !ok {
			c, ok = p.children[tgt[i]]
			p = p.failed
		}
		if p == nil {
			p = ac.root
		} else {
			p = c
		}
		if p.isTerm {
			fmt.Println(p.term, i-p.len+1, i)
		}
	}
}

func dfsAC(node *NodeAC) {
	visitedAC[node] = true
	for _, n := range node.children {
		if !visitedAC[n] {
			fmt.Println(string(n.val), n.id, n.len, n.isTerm, n.failed.id)
			dfsAC(n)
		}
	}
}

var visitedAC = make(map[*NodeAC]bool)
