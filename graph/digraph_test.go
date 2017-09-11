package digraph

import "testing"

func TestBasicOperation(t *testing.T) {
	g := NewGraph()
	u := NewNode([]byte("Wikipedia"))
	v := NewNode([]byte("Org"))

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

	if len(g.Nodes()) > 1 {
		t.Errorf("Expected single node to be in the graph.")
	}

	g.AddNodes(u, v)
	if !g.Has(u) {
		t.Errorf("Expected to have %s [node] in graph. ", u.ToString())
	}

	if !g.Has(v) {
		t.Errorf("Expected to have %s [node] in graph. ", v.ToString())
	}

	if !g.EdgeExist(u, v) {
		t.Errorf("Expected to have Edge exist between `%s` and `%s`", u, v)
	}

	x := NewNode([]byte("Outlier"))
	if g.Has(x) {
		t.Errorf("Expected not to have `%s` node.", x.ToString())
	}

	if g.EdgeExist(u, x) {
		t.Errorf("Expected not to have Edge exist between `%s` and `%s`", u, x)
	}
}

func getTestData() (g *Graph, u, v *Node) {
	g = NewGraph()
	u = NewNode([]byte("Wikipedia"))
	v = NewNode([]byte("Article"))
	return
}
