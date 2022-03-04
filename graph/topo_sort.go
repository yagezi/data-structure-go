/*
*@Time        : 2022/01/20
*@File        : topo_sort.go
*@Description : topological sort
 */

package graph

import (
	"container/list"
	"fmt"
)

// topoSort implements topological sort by kahn algorithm
func topoSort(g *Graph) {
	inDegree := make([]int, g.v)
	// alternative is to use stack
	queue := list.New()
	// count vertex num
	var count int

	// 1. count in-degree
	for _, list := range g.adj {
		for e := list.Front(); e != nil; e = e.Next() {
			edge := e.Value.(*Edge)
			inDegree[edge.tid]++
		}
	}

	// 2. find the vertex that in-degree = 0
	for i, v := range inDegree {
		if v == 0 {
			queue.PushBack(i)
		}
	}

	// 3. remove this vertex and its edges. decrease in-degree of another vertexes.
	// Then push the vertex that in-degree = 0 iteratively
	for queue.Len() > 0 {
		v := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		count++
		fmt.Printf("%d ", v)
		for e := g.adj[v].Front(); e != nil; e = e.Next() {
			edge := e.Value.(*Edge)
			inDegree[edge.tid]--
			if inDegree[edge.tid] == 0 {
				queue.PushBack(edge.tid)
			}
		}
	}

	if count < g.v {
		fmt.Println("circle exist")
	}
}
