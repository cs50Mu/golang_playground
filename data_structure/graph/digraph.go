package graph

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

// Grapher interface for Graph
type Grapher interface {
	AddEdge(v, w int)
}

// Digraph 有向图
type Digraph struct {
	// number of vertices
	V int
	// number of edges
	E int
	// 邻接表
	adj [][]int
}

func NewDigraph(v int) Grapher {
	adj := make([][]int, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]int, 0)
	}
	return &Digraph{
		V:   v,
		adj: adj,
	}
}

func (g *Digraph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.E++
}

func (g *Digraph) Adj(v int) []int {
	return g.adj[v]
}

// DirectedDFS 可达性
type DirectedDFS struct {
	*Digraph
	marked []bool
}

//
func NewDirectedDFS(g *Digraph, s int) *DirectedDFS {
	dd := &DirectedDFS{
		Digraph: g,
		marked:  make([]bool, g.V),
	}
	dd.dfs(s)

	return dd
}

func (dd *DirectedDFS) dfs(v int) {
	dd.marked[v] = true
	for _, vert := range dd.Adj(v) {
		if !dd.marked[vert] {
			dd.dfs(vert)
		}
	}
}

func (dd *DirectedDFS) Marked(v int) bool {
	return dd.marked[v]
}

// DirectedCycle 寻找有向环
type DirectedCycle struct {
	*Digraph
	marked  []bool
	edgeTo  []int
	cycle   *stack.Stack
	onStack []bool
}

func NewDirectedCycle(g *Digraph) *DirectedCycle {
	dc := &DirectedCycle{
		Digraph: g,
		marked:  make([]bool, g.V),
		edgeTo:  make([]int, g.V),
		onStack: make([]bool, g.V),
		cycle:   stack.New(),
	}
	// dfs every verticle
	for i := 0; i < dc.V; i++ {
		if !dc.marked[i] {
			dc.dfs(i)
		}
	}

	return dc
}

func (dc *DirectedCycle) dfs(v int) {
	dc.marked[v] = true
	dc.onStack[v] = true
	for _, vert := range dc.Adj(v) {
		// 若已找到一个环，则无需再找，立即退出
		if dc.HasCycle() {
			return
		}
		if !dc.marked[vert] {
			dc.edgeTo[vert] = v
			dc.dfs(vert)
		} else {
			// 如果 vert 之前已经访问过了而且此时 vert 还在栈上
			// 则一定存在环
			if dc.onStack[vert] {
				// record the cycle
				// 从 v 往前找 vert
				for x := v; x != vert; x = dc.edgeTo[x] {
					dc.cycle.Push(x)
				}
				// 再把现在的 v 和 vert 加上
				dc.cycle.Push(vert)
				dc.cycle.Push(v)
			}
		}
	}

	dc.onStack[v] = false
}

func (dc *DirectedCycle) HasCycle() bool {
	return dc.cycle.Len() > 0
}

func (dc *DirectedCycle) Cycle() []int {
	fmt.Printf("stack: %+v\n", *dc.cycle)
	var res []int
	for dc.cycle.Len() > 0 {
		tmp := dc.cycle.Pop()
		val := tmp.(int)
		res = append(res, val)
	}
	return res
}

// Topological 拓朴排序
type Topological struct {
	*Digraph
	cycleFinder *DirectedCycle
	reversePost *stack.Stack
	marked      []bool
}

func NewTopological(g *Digraph) *Topological {
	t := &Topological{
		Digraph:     g,
		cycleFinder: NewDirectedCycle(g),
		reversePost: stack.New(),
		marked:      make([]bool, g.V),
	}

	if !t.cycleFinder.HasCycle() {
		for i := 0; i < g.V; i++ {
			t.dfs(i)
		}
	}

	return t
}

func (t *Topological) dfs(v int) {
	t.marked[v] = true
	for _, vert := range t.Adj(v) {
		if !t.marked[vert] {
			t.dfs(vert)
		}
	}
	t.reversePost.Push(v)
}

func (t *Topological) IsDAG() bool {
	return t.reversePost.Len() > 0
}

func (t *Topological) Order() []int {
	var res []int
	for t.reversePost.Len() > 0 {
		tmp := t.reversePost.Pop()
		val := tmp.(int)
		res = append(res, val)
	}

	return res
}
