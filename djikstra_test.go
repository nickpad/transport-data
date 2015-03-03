package transportdata

import (
	"fmt"
	"testing"
)

func setupDjikstraGraph() (dGraph, *dVertex) {
	graph := make(dGraph)
	graph.connectVertices("milsons point", "wynyard")
	graph.connectVertices("wynyard", "town hall")
	graph.connectVertices("town hall", "central")
	graph.connectVertices("town hall", "martin place")
	graph.connectVertices("martin place", "kings cross")
	graph.connectVertices("kings cross", "edgecliff")
	graph.connectVertices("edgecliff", "bondi")
	graph.connectVertices("central", "museum")
	graph.connectVertices("museum", "st james")
	graph.connectVertices("st james", "circular quay")
	graph.connectVertices("circular quay", "wynyard")
	start := graph["milsons point"]
	return graph, start
}

func TestDjikstra(t *testing.T) {
	graph, start := setupDjikstraGraph()
	djikstra(graph, start)
	for _, vertex := range graph {
		fmt.Println(vertex)
	}
}

func BenchmarkDjikstra(b *testing.B) {
	graph, start := setupDjikstraGraph()

	for n := 0; n < b.N; n++ {
		djikstra(graph, start)
	}
}
