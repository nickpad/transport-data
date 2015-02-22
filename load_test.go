package transportdata

import "testing"

func TestLoadStops(t *testing.T) {
	graph := NewGraph()
	LoadStops("data/stops.txt", graph)
	stopId := "20611"
	expectedName := "Milsons Point Wharf"

	actualName := graph.nodes[stopId].Name

	if expectedName != actualName {
		t.Errorf("Unexpected stop name: %#v", actualName)
	}
}
