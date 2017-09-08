// Implementation of directed graph. This means A->B != B<-A
// Nodes <= ["doctor", "sujit", "ratna", "shakya"]
// Arcs <= {"doctor": ["sujit"], "sujit": ["ratna", "shakya"], "ratna": ["shakya"], "shakya": []}
package digraph

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

func (n *Node) ToString() string {
	return string(n.value)
}

func (g *Graph) Has(node *Node) bool {
	if _, ok:= g.arcs[node.value]; ok {
		return true
	}
	return false
}

func (g *Graph) AddNodes(from *Node, to *Node) {
	g.nodes = append(g.nodes, from)
	g.nodes = append(g.nodes, to)

	// Connect `from` node with `to` node.
	g.arcs[from] = append(g.arcs[from], to)
}

func (g *Graph) Nodes() []*Node {
	return g.nodes
}

func (g *Graph) Edges(from *Node) []*Node {
	return g.Arcs[from]
}
