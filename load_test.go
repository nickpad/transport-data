package transportdata

import "testing"

func TestLoadStops(t *testing.T) {
	graph := make(graph)
	stopName := "Milsons Point Wharf"

	loadStops("fixtures/stops.txt", graph)

	vertexId := graph[stopName].vertexId
	if vertexId != stopName {
		t.Errorf("Unexpected stop name: %#v", vertexId)
	}
}
