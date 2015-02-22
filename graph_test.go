package transportdata

import "testing"

func TestShortestPathSearch(t *testing.T) {
	graph := NewGraph()
	graph.AddStop("1", "Central")
	graph.AddStop("2", "Wynyard")
	graph.AddStop("3", "Town Hall")
	graph.AddStop("4", "Milsons Point")
	graph.AddEdge("1", "2")
	graph.AddEdge("2", "3")
	graph.AddEdge("3", "4")
	graph.AddEdge("1", "3")

	distance := graph.PathSearch("1", "4")

	if distance != 2 {
		t.Fatalf("Expected 2 but got %#v", distance)
	}
}
