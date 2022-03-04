/*
*@Time        : 2022/01/27
*@File        : utils.go
*@Description :
 */

package graph

type VertexHeap []*Vertex

func (h VertexHeap) Len() int            { return len(h) }
func (h VertexHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h VertexHeap) Less(i, j int) bool  { return h[i].value.(int) < h[j].value.(int) }
func (h *VertexHeap) Push(x interface{}) { *h = append(*h, x.(*Vertex)) }
func (h *VertexHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
