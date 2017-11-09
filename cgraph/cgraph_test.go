package cgraph

import "testing"

func TestConcurrentSafety(t *testing.T) {
	g := NewGraph()
	u := NewNode("one")
	v := NewNode("hundred")
	y := NewNode("thousand")
	g.AddNodes(u, v)
	g.AddNodes(u, y)

	if ok, _ := g.Has(u); ok {
		g.Edges(u)
	}
}
