package digraph

import "testing"

func setupTest() (g *Graph, u, v *Node) {
	g = NewGraph()
	u = NewNode([]byte("Wikipedia"))
	v = NewNode([]byte("Article"))
	return
}

func TestBasicOperation(t *testing.T) {
	g,u,v := setupTest()

	t.Log("Test: Add Single Node")
	g.AddNodes(u, nil)
	if !g.Has(u) {
		t.Errorf("Expected to have %s [node] in graph. ", u.ToString())
	}

	check := g.EdgeExist(u, nil)
	if check == true {
		t.Errorf("Expected not to have edge exist for `%s` and nil object", u.ToString())
	}

	check = g.EdgeExist(v, u)
	if check == true {
		t.Errorf("Expected not to have edge exist for `%s` and `%s`", v.ToString(), u.ToString())
	}

	check = g.EdgeExist(u, v)
	if check == true {
		t.Errorf("Expected not to have edge exist for `%s` and `%s`", u.ToString(), v.ToString())
	}

	if _,l:=g.Nodes(); l > 1 {
		t.Errorf("Expected single node to be in the graph.")
	}

	t.Log("Test: Add Nodes with edge.")
	g.AddNodes(u, v)
	if !g.Has(u) {
		t.Errorf("Expected to have %s [node] in graph. ", u.ToString())
	}

	if !g.Has(v) {
		t.Errorf("Expected to have %s [node] in graph. ", v.ToString())
	}

	if !g.EdgeExist(u, v) {
		t.Errorf("Expected to have Edge exist between `%s` and `%s`", u.ToString(), v.ToString())
	}

	x := NewNode([]byte("Outlier"))
	if g.Has(x) {
		t.Errorf("Expected not to have `%s` node.", x.ToString())
	}

	if g.EdgeExist(u, x) {
		t.Errorf("Expected not to have Edge exist between `%s` and `%s`", u.ToString(), x.ToString())
	}

	t.Log("Test: single node connected with multiple nodes")
	g.AddNodes(u, x)
	if _, l := g.Edges(u); l != 2 {
		t.Errorf("Expected to have two nodes connected with %s", u.ToString())
	}

	t.Log("Test: Chaining of node. x<-u->v->x")
	g.AddNodes(v, x)
	if _, l := g.Nodes(); l!=3 {
		t.Errorf("Expected to have two nodes in a graph")
	}
}

func BenchmarkAddSingleNode(b *testing.B) {
	g, u, _ := setupTest()
	for i:=0; i<b.N; i++ {
		g.AddNodes(u, nil)
	}
}

func BenchmarkAddNodeWithEdge(b *testing.B) {
	g, u, v := setupTest()
	for i:=0; i<b.N; i++ {
		g.AddNodes(u, v)
	}
}

func BenchmarkNodeConnectedWithTwoNode(b *testing.B) {
	g, u, v := setupTest()
	x := NewNode([]byte("X"))
	for i:=0; i<b.N; i++ {
		g.AddNodes(u, v)
		g.AddNodes(u, x)
	}
}

func BenchmarkAddChainOfNodes(b *testing.B) {
	g, u, v := setupTest()
	x := NewNode([]byte("Chaining"))
	for i:=0; i<b.N; i++ {
		g.AddNodes(u,v)
		g.AddNodes(v,x)
	}
}
