package toposort

import "fmt"

type Graph struct {
	nodes   []string
	outputs map[string]map[string]int
	inputs  map[string]int
}

func NewGraph(cap int) *Graph {
	return &Graph{
		nodes:   make([]string, 0, cap),
		inputs:  make(map[string]int),
		outputs: make(map[string]map[string]int),
	}
}

func (g *Graph) AddNode(name string) {
	g.nodes = append(g.nodes, name)
	g.outputs[name] = make(map[string]int)
	g.inputs[name] = 0
}

func (g *Graph) AddNodes(names ...string) {
	for _, name := range names {
		g.AddNode(name)
	}
}

func (g *Graph) AddEdge(from, to string) error {
	m, ok := g.outputs[from]
	if !ok {
		return fmt.Errorf("no such node: %v", from)
	}

	m[to] = len(m) + 1
	g.inputs[to]++

	return nil
}

func (g *Graph) RemoveEdge(from, to string) {
	delete(g.outputs[from], to)
	g.inputs[to]--
}

func (g *Graph) Toposort() ([]string, bool) {
	L := make([]string, 0, len(g.nodes))

	S := make([]string, 0)
	for _, n := range g.nodes {
		if g.inputs[n] == 0 {
			S = append(S, n)
		}
	}

	for len(S) > 0 {
		var n string
		n, S = S[0], S[1:]
		L = append(L, n)

		ms := make([]string, len(g.outputs[n]))
		for m, i := range g.outputs[n] {
			ms[i-1] = m
		}

		for _, m := range ms {
			g.RemoveEdge(n, m)

			if g.inputs[m] == 0 {
				S = append(S, m)
			}
		}
	}

	N := 0
	for _, v := range g.inputs {
		N += v
	}

	if N > 0 {
		return L, false
	}

	return L, true
}
