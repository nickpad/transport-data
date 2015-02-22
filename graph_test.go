package transportdata

import "testing"

func setupGraph() Graph {
	graph := NewGraph()
	graph.AddStop("1", "Central")
	graph.AddStop("2", "Wynyard")
	graph.AddStop("3", "Town Hall")
	graph.AddStop("4", "Milsons Point")
	graph.AddEdge("1", "2")
	graph.AddEdge("2", "3")
	graph.AddEdge("3", "4")
	graph.AddEdge("1", "3")

	return graph
}

func TestPathSearch(t *testing.T) {
	graph := setupGraph()
	distance := graph.PathSearch("1", "4")
	expected := 2

	if distance != expected {
		t.Fatalf("Expected %#v but got %#v", expected, distance)
	}
}

func BenchmarkPathSearch(b *testing.B) {
	graph := setupGraph()

	for n := 0; n < b.N; n++ {
		graph.PathSearch("1", "4")
	}
}
