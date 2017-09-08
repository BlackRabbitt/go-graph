package graph

type Node interface{
	// Returns string value of Node
	ToString() string
}

type Graph interface {
	// Has return true if Node exist
	Has(Node) bool

	// Add and connect two nodes.
	AddNodes(from, to *Node)

	// Nodes returns all nodes from the graph
	Nodes() []Node

	// Edges returns the edge nodes. If the return Array is empty, we can say, Edge doesnot exist beyond that.
	Edges(u Node) []*Node

	// EdgeExist returns true if two nodes are connected from u->v otherwise return false
	EdgeExist(u, v Node) bool
}
