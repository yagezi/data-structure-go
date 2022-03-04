/*
*@Time        : 2022/03/01
*@File        : lru.go
*@Description :
 */

package lru

import "container/list"

type LRU struct {
	cap  int
	maps map[int]*list.Element
	list *list.List
}

func New(cap int) *LRU {
	return &LRU{
		cap:  cap,
		maps: make(map[int]*list.Element),
		list: list.New(),
	}
}

func (l *LRU) Get(key int) int {
	if elem, ok := l.maps[key]; ok {
		l.list.MoveToBack(elem)
		return elem.Value.(int)
	}
	return -1
}

func (l *LRU) Put(key int) {
	if l.Get(key) == -1 {
		l.list.PushBack(key)
		l.maps[key] = l.list.Back()
	}
	if l.list.Len() > l.cap {
		elem := l.list.Front()
		l.list.Remove(elem)
		delete(l.maps, elem.Value.(int))
	}
}
