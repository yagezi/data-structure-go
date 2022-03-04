/*
*@Time        : 2022/01/27
*@File        : shortest_path.go
*@Description : shortest path algorithm
 */

package graph

import (
	"container/heap"
	"fmt"
)

// Dijkstra implements dijkstra's shortest path alogorithm,
// for directed or undirected graph with positive weight
func Dijkstra(g *Graph, src, dst int) {
	// using min-heap as priority queue, Vertex.value is distance to src
	h := &VertexHeap{}
	heap.Init(h)
	// initial distance array
	const MAX_DISTANCE = 65535
	distance := make([]int, g.v)
	for i := range distance {
		distance[i] = MAX_DISTANCE
	}
	// record path by path array
	path := make([]int, g.v)
	path[src] = src
	// initial src
	distance[src] = 0
	heap.Push(h, &Vertex{id: src, value: 0})
	// main loop
	for h.Len() > 0 {
		v := heap.Pop(h).(*Vertex)
		if v.value.(int) > distance[v.id] {
			continue
		}
		for e := g.adj[v.id].Front(); e != nil; e = e.Next() {
			edge := e.Value.(*Edge)
			if distance[edge.tid] > distance[v.id]+edge.w {
				path[edge.tid] = v.id
				distance[edge.tid] = distance[v.id] + edge.w
				heap.Push(h, &Vertex{id: edge.tid, value: distance[edge.tid]})
			}
		}

	}
	fmt.Printf("\ndistance:\n")
	for i, v := range distance {
		fmt.Println(i, v)
	}
	fmt.Printf("\npath:\n")
	for i, v := range path {
		fmt.Println(i, v)
	}
}

// BellmanFord implements Bellman-Ford's shortest path alogorithm,
// for directed graph with negetive weight
func BellmanFord(g *Graph, src, dst int) {
	const MAX_DISTANCE = 65535
	distance := make([]int, g.v)
	for i := range distance {
		distance[i] = MAX_DISTANCE
	}
	distance[src] = 0
	path := make([]int, g.v)
	path[src] = src
	for _, v := range g.adj {
		for e := v.Front(); e != nil; e = e.Next() {
			edge := e.Value.(*Edge)
			nu, nv := edge.sid, edge.tid
			if distance[nv] > distance[nu]+edge.w {
				distance[nv] = distance[nu] + edge.w
				path[nv] = nu
			}
		}
	}
	fmt.Printf("\ndistance:\n")
	for i, v := range distance {
		fmt.Println(i, v)
	}
	fmt.Printf("\npath:\n")
	for i, v := range path {
		fmt.Println(i, v)
	}
}
