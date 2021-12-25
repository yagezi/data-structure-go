/*
*@Time        : 2021/12/17
*@Creator     :
*@File        : skiplist.go
*@Description : 数据结构：跳表
 */

package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	MAX_LEVEL = 10
)

type Node struct {
	key   interface{}
	value interface{}
	next  []*Node
}

type ISkipList interface {
	Get(key interface{}) interface{}
	Insert(key, value interface{})
	Remove(key interface{}) bool
}

type SkipList struct {
	random        *rand.Rand
	head          Node
	len           int
	maxLevel      int
	cmp           CmpFunc
	preNodesCache []*Node

	isPrintDetail bool
}

func NewSkipList(opt ...OptionFunc) *SkipList {
	option := NewOption(opt...)
	sl := &SkipList{
		random:        option.random,
		maxLevel:      option.maxLevel,
		cmp:           option.cmp,
		isPrintDetail: option.isPrintDetail,
		len:           0,
		preNodesCache: make([]*Node, option.maxLevel),
		head: Node{
			key:   nil,
			value: nil,
			next:  make([]*Node, option.maxLevel),
		},
	}

	return sl
}

func (sl SkipList) randomLevel() int {
	level := 1
	// 提升概率为 0.5
	for rand.Int63() < (math.MaxInt64>>1) && level < sl.maxLevel {
		level += 1
	}
	return level
}

func (sl SkipList) findPreNodes(key interface{}) []*Node {
	// 利用缓存池preNodesCache，避免频繁开辟和回收空间
	preNodes := sl.preNodesCache
	pre := &sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		if pre != nil {
			for cur := pre.next[i]; cur != nil; cur = cur.next[i] {
				if sl.cmp(cur.key, key) >= 0 {
					break
				}
				pre = cur
			}

		}
		preNodes[i] = pre
	}
	return preNodes
}

func (sl SkipList) String() string {
	str := fmt.Sprintf("{{max level:%d}.{current len:%d}}\n", sl.maxLevel, sl.len)
	heads := sl.head.next
	for i := sl.maxLevel - 1; i >= 0; i-- {
		count := 0
		for cur := heads[i]; cur != nil; cur = cur.next[i] {
			if sl.isPrintDetail {
				str = fmt.Sprintf("%s {{key:%v},{value:%v}}", str, cur.key, cur.value)
			}
			count++
		}
		str = fmt.Sprintf("%s\n%d\n\n", str, count)
	}
	return str
}

func (sl SkipList) Get(key interface{}) interface{} {
	var pre = &sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		for cur := pre.next[i]; cur != nil; cur = cur.next[i] {
			cmpRes := sl.cmp(cur.key, key)
			if cmpRes == 0 {
				return cur.value
			}
			if cmpRes > 0 {
				break
			}
			pre = cur
		}
	}
	return nil
}

func (sl *SkipList) Insert(key, value interface{}) {
	preNodes := sl.findPreNodes(key)
	if preNodes[0].next[0] != nil && sl.cmp(preNodes[0].next[0].key, key) == 0 {
		// 最低层包含所有数据，其它层都是索引。索引层只有key有用，value是无用的。
		// key相同，需要更新value时，只需要更新最低层。
		preNodes[0].next[0].value = value
		return
	}

	level := sl.randomLevel()
	node := &Node{
		key:   key,
		value: value,
		next:  make([]*Node, level),
	}
	for i := range node.next {
		node.next[i] = preNodes[i].next[i]
		preNodes[i].next[i] = node
	}
	sl.len++
}

func (sl *SkipList) Remove(key interface{}) bool {
	preNodes := sl.findPreNodes(key)
	cur := preNodes[0].next[0]
	if cur == nil {
		return false
	} else if sl.cmp(cur.key, key) != 0 {
		return false
	}

	for i := range cur.next {
		preNodes[i].next[i] = cur.next[i]
	}
	return true
}
