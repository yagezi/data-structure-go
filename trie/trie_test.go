package trie

import (
	"fmt"
	"testing"
)

func TestMask(t *testing.T) {
	rs := []rune{'a', 'b', 'c', 'd'}
	var m uint64
	for _, r := range rs {
		fmt.Printf("%x\n", uint64(r))
		m |= uint64(1) << uint64(r-'0')
		// fmt.Println(m)
	}
}

func TestTrie(t *testing.T) {
	tree := NewTrie()
	// tree.Add("foo")
	tree.Add("bcd")
	tree.Add("faobar")
	tree.Add("faolish")
	fmt.Println(tree)

	fmt.Println(tree.HasKeyWithPrefix("fao"))
	fmt.Println(tree.HasKeyWithPrefix("bar"))

	fmt.Println(tree.Search("fao"))

}
