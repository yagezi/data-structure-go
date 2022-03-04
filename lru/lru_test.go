/*
*@Time        : 2022/03/01
*@File        : lru_test.go
*@Description :
 */

package lru

import (
	"fmt"
	"testing"
)

func Test_LRU(t *testing.T) {
	lru := New(4)
	lru.Put('a')
	lru.Put('b')
	lru.Put('c')
	for k, v := range lru.maps {
		fmt.Println(k, v.Value.(int))
	}
	for e := lru.list.Front(); e != nil; e = e.Next() {
		fmt.Printf(" %d", e.Value.(int))
	}
	fmt.Println("")

	fmt.Println(lru.Get('a'))
	// fmt.Println(lru.Get('d'))
	for e := lru.list.Front(); e != nil; e = e.Next() {
		fmt.Printf(" %d", e.Value.(int))
	}
	fmt.Println("")

	lru.Put('c')
	for e := lru.list.Front(); e != nil; e = e.Next() {
		fmt.Printf(" %d", e.Value.(int))
	}
	fmt.Println("")
}
