/*
*@Time        : 2022/01/12
*@File        : mst.go
*@Description : minimum spanning tree
 */

package graph

import (
	"fmt"
	"sort"

	"github.com/yagezi/data-structure-go/dsu"
)

// TODO: prim and kruskal

func mst_kruskal(g *Graph) *Graph {
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
	res := NewGraph(g.v)
	for _, e := range edges {
		if !dsu.In(e.sid, e.tid) {
			dsu.Union(e.sid, e.tid)
			res.addEdge(e.sid, e.tid, e.w)
			if res.e == res.v-1 {
				break
			}
		}
	}
	return res
}

func dbgPrintEdges(edges []*Edge) {
	fmt.Println("sorted edges:")
	for _, v := range edges {
		fmt.Println(v)
	}
}
