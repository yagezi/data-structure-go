/*
*@Time        : 2022/01/14
*@File        : dsu.go
*@Description : disjoint set union, a.k.a. union-find set
 */

package dsu

import "fmt"

type DSU struct {
	parents []int
}

func NewDSU(n int) *DSU {
	d := &DSU{parents: make([]int, n)}
	for i := range d.parents {
		d.parents[i] = i
	}
	return d
}

func (d *DSU) Find(x int) int {
	for d.parents[x] != x {
		x = d.parents[x]
	}
	return x
}

func (d *DSU) Union(x, y int) {
	root := d.Find(x)
	d.parents[d.Find(y)] = root
	// path compress
	var i, j int
	i = x
	for i != root {
		j = d.parents[i]
		d.parents[i] = root
		i = j
	}
}

func (d *DSU) In(x, y int) bool {
	return d.Find(x) == d.Find(y)
}

func (d DSU) String() string {
	m := make(map[int][]int)
	for i, v := range d.parents {
		root := d.Find(v)
		if _, exist := m[root]; !exist {
			m[root] = make([]int, 0)
		}
		m[root] = append(m[root], i)
	}
	str := fmt.Sprintf("parents %v\n", d.parents)
	for _, v := range m {
		str = fmt.Sprintln(str, v)
	}
	return str
}
