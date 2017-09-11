// Implementation of directed graph. This means A->B != B<-A
package digraph

import "bytes"

// Node is a simple graph node.
type Node struct {
	value []byte
}

// Graph is a set of Nodes, Edges
type Graph struct {
	nodes []*Node
	arcs map[*Node][]*Node
}

func NewGraph() *Graph {
	graph := new(Graph)
	graph.arcs = make(map[*Node][]*Node)
	return graph
}

func NewNode(v []byte) *Node {
	return &Node{value: []byte(v)}
}

func (n *Node) ToString() string {
	return string(n.value)
}

func (g *Graph) Has(node *Node) bool {
	for i := range g.nodes {
		if g.nodes[i].equal(node) {
			return true
		}
	}
	return false
}

// if v is nil, then the graph will have single node with no arcs.
func (g *Graph) AddNodes(u *Node, v *Node) {
	if !g.Has(u) {
		g.nodes = append(g.nodes, u)
	}

	if v != nil {
		if !g.Has(v) {
			g.nodes = append(g.nodes, v)
		}

		// Connect `u` node with `v` node.
		if !g.EdgeExist(u, v) {
			g.arcs[u] = append(g.arcs[u], v)
		}
	}
}

// returns all nodes in the graph
func (g *Graph) Nodes() ([]*Node, int) {
	return g.nodes, len(g.nodes)
}

// EdgeExists returns true if atleast one arc exist from `u` to `v`
func (g *Graph) EdgeExist(u, v *Node) bool {
	if adjArcs, ok:= g.arcs[u]; ok {
		for i := range adjArcs {
			if adjArcs[i].equal(v) {
				return true
			}
		}
	}
	return false
}

// returns adjacent arc nodes for node `u`
func (g *Graph) Edges(u *Node) ([]*Node, int) {
	return g.arcs[u], len(g.arcs[u])
}

func (u *Node) equal(v *Node) bool {
	if bytes.Equal(u.value, v.value) {
		return true
	}
	return false
}
