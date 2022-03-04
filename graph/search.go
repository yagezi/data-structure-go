/*
*@Time        : 2022/01/12
*@File        : search.go
*@Description : graph search algorithm, just like bfs and dfs
 */

package graph

import (
	"container/list"
	"fmt"
)

var visited = make(map[int]bool)

func dfs(g *Graph, i int) {
	visited[i] = true
	fmt.Println(i)
	for e := g.adj[i].Front(); e != nil; e = e.Next() {
		j := e.Value.(*Edge).tid
		if !visited[j] {
			dfs(g, j)
			fmt.Printf("\n")
		}
	}
}

func dfsIter(g *Graph, i int) {
	stack := list.New()
	stack.PushBack(i)
	for stack.Len() > 0 {
		v := stack.Back().Value.(int)
		stack.Remove(stack.Back())
		if !visited[v] {
			fmt.Println(v)
			visited[v] = true
			for e := g.adj[v].Front(); e != nil; e = e.Next() {
				j := e.Value.(*Edge).tid
				if !visited[j] {
					stack.PushBack(j)
				}
			}
		}
	}

}

// 对非连通图进行遍历
func dfsTraverse(g *Graph) {
	for i := 0; i < g.v; i++ {
		if !visited[i] {
			dfs(g, i)
		}
	}
}

func bfs(g *Graph, i int) {
	queue := list.New()
	queue.PushBack(i)
	for queue.Len() > 0 {
		v := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		if !visited[v] {
			fmt.Println(v)
			visited[v] = true
			for e := g.adj[v].Front(); e != nil; e = e.Next() {
				j := e.Value.(*Edge).tid
				if !visited[j] {
					queue.PushBack(j)
				}
			}
		}
	}
}
