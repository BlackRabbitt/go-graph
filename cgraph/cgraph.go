// Digraph safe for concurrent use.

package cgraph

import (
	"bytes"
	"strings"

	"sync"
)

// Node is a simple graph node.
type Node struct {
	value []byte
}

// Graph is a set of Nodes, Edges
type Graph struct {
	nodes []*Node

	// protect maps from concurrent read/write
	sync.RWMutex

	arcs map[*Node][]*Node
}

func NewGraph() *Graph {
	graph := new(Graph)
	graph.arcs = make(map[*Node][]*Node)
	return graph
}

func NewNode(v string) *Node {
	return &Node{value: []byte(strings.ToLower(v))}
}

func (n *Node) ToString() string {
	return string(n.value)
}

func (g *Graph) Has(node *Node) (bool, *Node) {
	for i := range g.nodes {
		if node.equal(g.nodes[i]) {
			return true, g.nodes[i]
		}
	}
	return false, node
}

// if v is nil, then the graph will have single node with no arcs.
func (g *Graph) AddNodes(u *Node, v *Node) {
	var ok bool
	ok, u = g.Has(u)
	if !ok {
		g.nodes = append(g.nodes, u)
	}

	if len(v.value)>0 {
		ok, v = g.Has(v)
		if !ok {
			g.nodes = append(g.nodes, v)
		}

		// Connect `u` node with `v` node.
		if !g.EdgeExist(u, v) {
			g.Lock()
			g.arcs[u] = append(g.arcs[u], v)
			g.Unlock()
		}
	}
}

// returns all nodes in the graph
func (g *Graph) Nodes() ([]*Node, int) {
	return g.nodes, len(g.nodes)
}

// EdgeExists returns true if atleast one arc exist from `u` to `v`
func (g *Graph) EdgeExist(u, v *Node) bool {
	_, u = g.Has(u)

	g.RLock()
	adjArcs, ok:= g.arcs[u]
	g.RUnlock()
	if ok {
		for i := range adjArcs {
			if v.equal(adjArcs[i]) {
				return true
			}
		}
	}
	return false
}
// returns adjacent arc nodes for node `u`
func (g *Graph) Edges(u *Node) ([]*Node, int) {
	_, u = g.Has(u)

	g.RLock()
	edges := g.arcs[u]
	g.RUnlock()
	return edges, len(edges)
}

// returns true if edge has arcs.
func (g *Graph) HasEdges(u *Node) bool {
	_, l := g.Edges(u)
	if l>0 {
		return true
	}
	return false
}

func (u *Node) equal(v *Node) bool {
	if bytes.Equal(u.value, v.value) {
		return true
	}
	return false
}
