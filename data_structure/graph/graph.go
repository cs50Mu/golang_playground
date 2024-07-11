package graph

import (
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

type Graph struct {
	// number of vertices
	V int
	// number of edges
	E int
	// 邻接表
	adj [][]int
}

func NewGraph(v int) Grapher {
	adj := make([][]int, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]int, 0)
	}
	return &Graph{
		V:   v,
		adj: adj,
	}
}

func (g *Graph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.E++
}

func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}

type Paths struct {
	*Graph
	marked []bool
	edgeTo []int
	// starting point
	s int
}

func newPaths(g *Graph, s int) *Paths {
	return &Paths{
		Graph:  g,
		marked: make([]bool, g.V),
		edgeTo: make([]int, g.V),
		s:      s,
	}
}

func NewDfsPaths(g *Graph, s int) *Paths {
	dp := newPaths(g, s)
	dp.dfs(s)

	return dp
}

func (dp *Paths) dfs(v int) {
	dp.marked[v] = true
	for _, vert := range dp.Adj(v) {
		if !dp.marked[vert] {
			//
			dp.edgeTo[vert] = v
			dp.dfs(vert)
		}
	}
}

func NewBfsPaths(g *Graph, s int) *Paths {
	dp := newPaths(g, s)
	dp.bfs(s)

	return dp
}

func (dp *Paths) bfs(v int) {
	q := queue.New()
	q.Enqueue(v)
	dp.marked[v] = true
	for q.Len() > 0 {
		tmp := q.Dequeue()
		val, _ := tmp.(int)
		for _, vert := range dp.Adj(val) {
			if !dp.marked[vert] {
				dp.edgeTo[vert] = val
				dp.marked[vert] = true
				q.Enqueue(vert)
			}
		}
	}
}

func (dp *Paths) HasPathTo(s, v int) bool {
	return dp.marked[v]
}

func (dp *Paths) PathTo(v int) []int {
	if !dp.marked[v] {
		return nil
	}

	stack := stack.New()
	// 依次 push 进了 v ... s
	for x := v; x != dp.s; x = dp.edgeTo[x] {
		stack.Push(x)
	}
	stack.Push(dp.s)

	var res []int
	// 依次 pop 出了 s ... v
	for stack.Len() > 0 {
		tmp := stack.Pop()
		val, _ := tmp.(int)
		res = append(res, val)
	}
	return res
}
