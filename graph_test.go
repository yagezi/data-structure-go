/*
*@Time        : 2022/01/12
*@File        : graph_test.go
*@Description :
 */
package graph

import (
	"fmt"
	"testing"
)

func Test_GraphCreate(t *testing.T) {
	g := NewGraph(3)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	fmt.Println(g)
}

func Test_mstKruskal(t *testing.T) {
	g := NewGraph(9)
	g.addEdge(0, 1, 10)
	g.addEdge(0, 5, 11)
	g.addEdge(1, 2, 18)
	g.addEdge(1, 6, 16)
	g.addEdge(1, 8, 12)
	g.addEdge(2, 3, 22)
	g.addEdge(2, 8, 8)
	g.addEdge(3, 4, 20)
	g.addEdge(3, 6, 24)
	g.addEdge(3, 7, 16)
	g.addEdge(3, 8, 21)
	g.addEdge(4, 5, 26)
	g.addEdge(4, 7, 7)
	g.addEdge(5, 6, 17)
	g.addEdge(6, 7, 19)

	mst := mst_kruskal(g)
	fmt.Println(mst)
	fmt.Println("sum of weight:", mst.sumWeight())
}
