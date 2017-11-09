// Digraph safe for concurrent use.

package cgraph

import (
	"bytes"
	"strings"

	"sync"
)

// Node is a simple graph node.
type Node struct {
	Value []byte
}

// Graph is a set of Nodes, Edges
type Graph struct {
	Nodes []*Node

	// protect maps from concurrent read/write
	sync.RWMutex

	Arcs map[*Node][]*Node
}

func NewGraph() *Graph {
	graph := new(Graph)
	graph.Arcs = make(map[*Node][]*Node)
	return graph
}

func NewNode(v string) *Node {
	return &Node{Value: []byte(strings.ToLower(v))}
}

func (n *Node) ToString() string {
	return string(n.Value)
}

func (g *Graph) Has(node *Node) (bool, *Node) {
	for i := range g.Nodes {
		if node.equal(g.Nodes[i]) {
			return true, g.Nodes[i]
		}
	}
	return false, node
}

// if v is nil, then the graph will have single node with no arcs.
func (g *Graph) AddNodes(u *Node, v *Node) {
	var ok bool
	ok, u = g.Has(u)
	if !ok {
		g.Nodes = append(g.Nodes, u)
	}

	if len(v.Value) > 0 {
		ok, v = g.Has(v)
		if !ok {
			g.Nodes = append(g.Nodes, v)
		}

		// Connect `u` node with `v` node.
		if !g.EdgeExist(u, v) {
			g.Lock()
			g.Arcs[u] = append(g.Arcs[u], v)
			g.Unlock()
		}
	}
}

// returns all nodes in the graph
func (g *Graph) AllNodes() ([]*Node, int) {
	return g.Nodes, len(g.Nodes)
}

// EdgeExists returns true if atleast one arc exist from `u` to `v`
func (g *Graph) EdgeExist(u, v *Node) bool {
	_, u = g.Has(u)

	g.RLock()
	adjArcs, ok := g.Arcs[u]
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
	edges := g.Arcs[u]
	g.RUnlock()
	return edges, len(edges)
}

// returns true if edge has arcs.
func (g *Graph) HasEdges(u *Node) bool {
	_, l := g.Edges(u)
	if l > 0 {
		return true
	}
	return false
}

func (u *Node) equal(v *Node) bool {
	if bytes.Equal(u.Value, v.Value) {
		return true
	}
	return false
}
