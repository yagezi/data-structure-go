/*
*@Time        : 2022/01/17
*@File        : search_test.go
*@Description :
 */

package search

import (
	"fmt"
	"testing"
)

func Test_biSearch(t *testing.T) {
	list := []int{1, 2, 3}
	// fmt.Println(BinarySearchRecur(list, 4))
	fmt.Println(BinarySearchIter(list, 2))
}
