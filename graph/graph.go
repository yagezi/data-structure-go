/*
*@Time        : 2022/01/12
*@File        : graph.go
*@Description :
 */

package graph

import (
	"container/list"
	"fmt"
)

type Edge struct {
	sid int
	tid int
	w   int
}

type Vertex struct {
	value interface{}
}

type Graph struct {
	v   int
	e   int
	adj []*list.List
	ves []Vertex
}

func NewGraph(v int, opts ...OptionFunc) *Graph {
	o := NewOption(opts)
	g := &Graph{
		v:   v,
		e:   0,
		adj: make([]*list.List, v),
	}
	for i := 0; i < g.v; i++ {
		g.adj[i] = list.New()
	}
	if o.hasVertexValue {
		g.ves = make([]Vertex, v)
	}
	return g
}

func (g *Graph) addEdge(src, dst int, weight ...int) {
	w := 0
	if len(weight) > 0 {
		w = weight[0]
	}
	g.adj[src].PushBack(&Edge{sid: src, tid: dst, w: w})
	g.e++
}

func (g Graph) sumWeight() (sum int) {
	for _, l := range g.adj {
		for e := l.Front(); e != nil; e = e.Next() {
			sum += e.Value.(*Edge).w
		}
	}
	return
}

func (g Graph) String() string {
	ret := fmt.Sprintf("vertexex: %d\n", g.v)
	for i, v := range g.adj {
		ret = fmt.Sprintf("%s v %d: ", ret, i)
		for e := v.Front(); e != nil; e = e.Next() {
			ret = fmt.Sprintf("%s (%d, %d) ", ret, i, e.Value.(*Edge).tid)
		}
		ret = fmt.Sprintln(ret)
	}
	return ret
}

func (g *Graph) SetValue(vid int, val interface{}) {
	g.ves[vid].value = val
}

func (g Graph) Value(vid int) interface{} {
	return g.ves[vid].value
}
