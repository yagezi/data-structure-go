/*
*@Time        : 2022/01/12
*@File        : mst.go
*@Description : minimum spanning tree
 */

package graph

import (
	"container/heap"
	"fmt"
	"sort"

	"github.com/yagezi/data-structure-go/dsu"
)

// tips: using disjoint set union
func mst_kruskal(g *Graph) {
	weightSum := 0
	dsu := dsu.NewDSU(g.e)
	edges := make([]*Edge, 0, g.e)
	for _, es := range g.adj {
		for e := es.Front(); e != nil; e = e.Next() {
			edges = append(edges, e.Value.(*Edge))
		}
	}
	sort.Slice(edges, func(i int, j int) bool {
		return edges[i].w < edges[j].w
	})
	// dbgPrintEdges(edges)
	for _, e := range edges {
		if !dsu.In(e.sid, e.tid) {
			dsu.Union(e.sid, e.tid)
			weightSum += e.w
			fmt.Printf("src: %d, dst %d, weight: %d\n", e.sid, e.tid, e.w)
		}
	}
	fmt.Println("sum of weight:", weightSum)
}

func dbgPrintEdges(edges []*Edge) {
	fmt.Println("sorted edges:")
	for _, v := range edges {
		fmt.Println(v)
	}
}

// tips: using priority queue or min-heap
func mst_prim(g *Graph) {
	h := &VertexHeap{}
	heap.Init(h)
	mstSet := make([]bool, g.v)
	weightSum := 0
	// Vertex.value means distance to mst
	heap.Push(h, &Vertex{id: 0, value: 0})
	for h.Len() > 0 {
		v := heap.Pop(h).(*Vertex)
		if !mstSet[v.id] {
			mstSet[v.id] = true
			weightSum += v.value.(int)
			fmt.Printf("node id: %d, weight: %d\n", v.id, v.value)
			for e := g.adj[v.id].Front(); e != nil; e = e.Next() {
				edge := e.Value.(*Edge)
				if !mstSet[edge.tid] {
					heap.Push(h, &Vertex{id: edge.tid, value: edge.w})
				}
			}
		}
	}
	fmt.Println("sum of weight:", weightSum)
}

