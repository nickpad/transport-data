package transportdata

import "testing"

func TestSuccesfulPathSearch(t *testing.T) {
	graph := buildTestGraph()
	start := graph["milsons point"]
	end := graph["martin place"]
	expected := "milsons point -> wynyard -> town hall -> martin place"
	state := NewState(graph, start, end, 0)

	state.Search()
	path := state.Path()

	if path.String() != expected {
		t.Fatalf("Expected %v but got %v", expected, path)
	}
}

func TestImpossibleRoute(t *testing.T) {
	graph := buildTestGraph()
	graph.addVertex("nowhere")
	start := graph["milsons point"]
	end := graph["nowhere"]
	state := NewState(graph, start, end, 0)

	state.Search()
	path := state.Path()

	if path.String() != "" {
		t.Fatalf("Expected empty path but got %v", path)
	}
}

func TestDepartureTimeAllowsShorterRoute(t *testing.T) {
	state := buildDepartureTimeTestGraph("town hall", 0)
	expected := "milsons point -> town hall"

	state.Search()
	path := state.Path()

	if path.String() != expected {
		t.Fatalf("Expected %v but got %v", expected, path)
	}
}

func TestDepartureTimeRequiresLongerRoute(t *testing.T) {
	state := buildDepartureTimeTestGraph("town hall", 2)
	expected := "milsons point -> wynyard -> town hall"

	state.Search()
	path := state.Path()

	if path.String() != expected {
		t.Fatalf("Expected %v but got %v", expected, path)
	}
}

func TestDepartureTimeMeansNoRouteIsPossible(t *testing.T) {
	state := buildDepartureTimeTestGraph("town hall", 10)

	state.Search()
	path := state.Path()

	if path.String() != "" {
		t.Fatalf("Expected empty path but got %v", path)
	}
}

func BenchmarkDjikstra(b *testing.B) {
	graph := buildTestGraph()
	start := graph["milsons point"]
	end := graph["central"]

	for n := 0; n < b.N; n++ {
		state := NewState(graph, start, end, 0)
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

func buildDepartureTimeTestGraph(destination string, departAt int64) *State {
	graph := make(graph)
	addRoute(graph, []string{"milsons point", "town hall"}, 0)
	addRoute(graph, []string{"milsons point", "wynyard", "town hall"}, 2)
	start := graph["milsons point"]
	end := graph[destination]

	return NewState(graph, start, end, departAt)
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
