/*
*@Time        : 2022/01/14
*@File        : dsu_test.go
*@Description :
 */

package dsu

import (
	"fmt"
	"testing"
)

func Test_dsu(t *testing.T) {
	dsu := NewDSU(10)
	dsu.Union(0, 1)
	dsu.Union(2, 3)
	dsu.Union(2, 4)
	dsu.Union(4, 5)
	dsu.Union(5, 6)

	fmt.Println(dsu)
	fmt.Println(dsu.In(2, 4))
	fmt.Println(dsu.In(1, 4))
}
