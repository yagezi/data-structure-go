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

func Test_DFS(t *testing.T) {
	g := NewGraph(3)
	g.addEdge(0, 1)
	g.addEdge(1, 0)
	g.addEdge(0, 2)
	g.addEdge(2, 0)
	g.addEdge(1, 2)
	g.addEdge(2, 1)

	dfs(g, 0)
}

func Test_DFSTraverse(t *testing.T) {
	g := NewGraph(6)
	g.addEdge(0, 1)
	g.addEdge(1, 0)
	g.addEdge(0, 2)
	g.addEdge(2, 0)
	g.addEdge(1, 2)
	g.addEdge(2, 1)

	g.addEdge(3, 4)
	g.addEdge(4, 3)
	g.addEdge(3, 5)
	g.addEdge(5, 3)
	g.addEdge(4, 5)
	g.addEdge(5, 4)

	dfsTraverse(g)
}

func Test_DFSIteration(t *testing.T) {
	g := NewGraph(8)
	g.addEdge(0, 1)
	g.addEdge(1, 0)
	g.addEdge(0, 2)
	g.addEdge(2, 0)
	g.addEdge(1, 2)
	g.addEdge(2, 1)

	g.addEdge(3, 4)
	g.addEdge(4, 5)
	g.addEdge(3, 6)
	g.addEdge(6, 7)

	dfsIter(g, 0)
	fmt.Printf("\n")
	dfsIter(g, 3)
}

func Test_BFS(t *testing.T) {
	g := NewGraph(8)
	g.addEdge(0, 1)
	g.addEdge(1, 0)
	g.addEdge(0, 2)
	g.addEdge(2, 0)
	g.addEdge(1, 2)
	g.addEdge(2, 1)

	g.addEdge(3, 4)
	g.addEdge(4, 5)
	g.addEdge(3, 6)
	g.addEdge(6, 7)

	bfs(g, 0)
	fmt.Printf("\n")
	bfs(g, 3)
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

	mst_kruskal(g)
}

func Test_prims(t *testing.T) {
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
	// from dst to src
	g.addEdge(1, 0, 10)
	g.addEdge(5, 0, 11)
	g.addEdge(2, 1, 18)
	g.addEdge(6, 1, 16)
	g.addEdge(8, 1, 12)
	g.addEdge(3, 2, 22)
	g.addEdge(8, 2, 8)
	g.addEdge(4, 3, 20)
	g.addEdge(6, 3, 24)
	g.addEdge(7, 3, 16)
	g.addEdge(8, 3, 21)
	g.addEdge(5, 4, 26)
	g.addEdge(7, 4, 7)
	g.addEdge(6, 5, 17)
	g.addEdge(7, 6, 19)

	mst_prim(g)
}

func Test_topoSort(t *testing.T) {
	g := NewGraph(14)
	g.addEdge(0, 4)
	g.addEdge(0, 5)
	g.addEdge(0, 11)
	g.addEdge(1, 2)
	g.addEdge(1, 4)
	g.addEdge(1, 8)
	g.addEdge(2, 5)
	g.addEdge(2, 6)
	g.addEdge(2, 9)
	g.addEdge(3, 2)
	g.addEdge(3, 13)
	g.addEdge(4, 7)
	g.addEdge(5, 8)
	g.addEdge(5, 12)
	g.addEdge(6, 5)
	g.addEdge(8, 7)
	g.addEdge(9, 10)
	g.addEdge(9, 11)
	g.addEdge(10, 13)
	g.addEdge(12, 9)

	topoSort(g)
	fmt.Println()

	// the edge leading to circle
	g.addEdge(7, 0)

	topoSort(g)
}

func Test_dijkstra(t *testing.T) {
	g := NewGraph(6)
	// 0-1
	g.addEdge(0, 1, 5)
	g.addEdge(1, 0, 5)
	// 0-2
	g.addEdge(0, 2, 1)
	g.addEdge(2, 0, 1)
	// 1-2
	g.addEdge(1, 2, 3)
	g.addEdge(2, 1, 3)
	// 1-4
	g.addEdge(1, 4, 8)
	g.addEdge(4, 1, 8)
	// 2-3
	g.addEdge(2, 3, 2)
	g.addEdge(3, 2, 2)
	// 2-4
	g.addEdge(2, 4, 1)
	g.addEdge(4, 2, 1)
	// 3-4
	g.addEdge(3, 4, 2)
	g.addEdge(4, 3, 2)
	// 3-5
	g.addEdge(3, 5, 1)
	g.addEdge(5, 3, 1)
	// 4-5
	g.addEdge(4, 5, 3)
	g.addEdge(5, 4, 3)

	Dijkstra(g, 0, 5)

}

func Test_BellmanFord(t *testing.T) {
	g := NewGraph(3)
	g.addEdge(0, 1, 4)
	g.addEdge(0, 2, 3)
	g.addEdge(1, 2, -2)

	fmt.Println("dijkstra")
	Dijkstra(g, 0, 2)
	fmt.Println("bellman-ford")
	BellmanFord(g, 0, 2)
}
