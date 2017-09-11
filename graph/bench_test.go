package digraph

import "testing"

func BenchMarkOperation(b *testing.B) {
	g,u,v := getTestData()
	//b.Log("----Adding single node in a graph----")
	for i:=0; i<b.N; i++ {
		g.AddNodes(u, nil)
	}
	//b.Log("----Adding two nodes with arc in a graph----")
	for i:=0; i<b.N; i++ {
		g.AddNodes(u, v)
	}
}
