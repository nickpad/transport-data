package transportdata

import "testing"

func TestSuccesfulPathSearch(t *testing.T) {
	graph := buildTestGraph()
	start := graph["milsons point"]
	end := graph["martin place"]
	expected := "milsons point -> wynyard -> town hall -> martin place"
	state := NewState(graph, start, 0)

	state.Search()
	path := state.PathTo(end)

	if path.String() != expected {
		t.Fatalf("Expected %v but got %v", expected, path)
	}
}

func TestImpossibleRoute(t *testing.T) {
	graph := buildTestGraph()
	graph.addVertex("nowhere")
	start := graph["milsons point"]
	end := graph["nowhere"]
	state := NewState(graph, start, 0)

	state.Search()
	path := state.PathTo(end)

	if path.String() != "" {
		t.Fatalf("Expected empty path but got %v", path)
	}
}

func TestDepartureTimeAllowsShorterRoute(t *testing.T) {
	state := buildDepartureTimeTestGraph(0)
	destination := state.graph["town hall"]
	expected := "milsons point -> town hall"

	state.Search()
	path := state.PathTo(destination)

	if path.String() != expected {
		t.Fatalf("Expected %v but got %v", expected, path)
	}
}

func TestDepartureTimeRequiresLongerRoute(t *testing.T) {
	state := buildDepartureTimeTestGraph(2)
	destination := state.graph["town hall"]
	expected := "milsons point -> wynyard -> town hall"

	state.Search()
	path := state.PathTo(destination)

	if path.String() != expected {
		t.Fatalf("Expected %v but got %v", expected, path)
	}
}

func BenchmarkDjikstra(b *testing.B) {
	graph := buildTestGraph()
	start := graph["milsons point"]

	for n := 0; n < b.N; n++ {
		state := NewState(graph, start, 0)
		state.Search()
	}
}

func buildTestGraph() graph {
	graph := make(graph)
	addRoute(graph, []string{"milsons point", "wynyard", "town hall", "central"}, 0)
	addRoute(graph, []string{"town hall", "martin place", "kings cross", "edgecliff", "bondi"}, 4)
	addRoute(graph, []string{"central", "museum", "st james", "circular quay", "wynyard"}, 6)
	return graph
}

func buildDepartureTimeTestGraph(departAt int64) *State {
	graph := make(graph)
	addRoute(graph, []string{"milsons point", "town hall"}, 0)
	addRoute(graph, []string{"milsons point", "wynyard", "town hall"}, 2)
	start := graph["milsons point"]

	return NewState(graph, start, departAt)
}

func addRoute(g graph, stops []string, startTime int64) {
	time := startTime
	for _, name := range stops {
		g.addVertex(name)
	}
	for i := 0; i < len(stops)-1; i++ {
		vertex := g[stops[i]]
		nextVertex := g[stops[i+1]]
		vertex.AddEdge(nextVertex, time, time+1)
		time = time + 2
	}
}
