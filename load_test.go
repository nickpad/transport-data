package transportdata

import "testing"

func TestLoadStops(t *testing.T) {
	graph := make(graph)
	stopName := "Milsons Point Wharf"

	loadStops("fixtures/stops.txt", graph)

	vertexID := graph[stopName].VertexID
	if vertexID != stopName {
		t.Errorf("Unexpected stop name: %#v", vertexID)
	}
}
