package transportdata

import (
	"fmt"
	"testing"
)

func (graph graph) AddNodeWithNameAndId(name string) {
	graph.AddNode(name, name)
}

func (graph graph) ConnectNodes(from, to string) {
	graph.AddEdge(from, to)
	graph.AddEdge(to, from)
}

func setupGraph() graph {
	graph := NewGraph()

	graph.AddNodeWithNameAndId("milsons point")
	graph.AddNodeWithNameAndId("wynyard")
	graph.AddNodeWithNameAndId("town hall")
	graph.AddNodeWithNameAndId("central")
	graph.AddNodeWithNameAndId("museum")
	graph.AddNodeWithNameAndId("st james")
	graph.AddNodeWithNameAndId("circular quay")
	graph.AddNodeWithNameAndId("martin place")
	graph.AddNodeWithNameAndId("kings cross")
	graph.AddNodeWithNameAndId("edgecliff")
	graph.AddNodeWithNameAndId("bondi junction")

	graph.ConnectNodes("wynyard", "milsons point")
	graph.ConnectNodes("wynyard", "town hall")
	graph.ConnectNodes("town hall", "central")
	graph.ConnectNodes("central", "museum")
	graph.ConnectNodes("museum", "st james")
	graph.ConnectNodes("st james", "circular quay")
	graph.ConnectNodes("circular quay", "wynyard")
	graph.ConnectNodes("town hall", "martin place")
	graph.ConnectNodes("martin place", "kings cross")
	graph.ConnectNodes("kings cross", "edgecliff")
	graph.ConnectNodes("edgecliff", "bondi junction")

	return graph
}

func TestSuccesfulPathSearch(t *testing.T) {
	graph := setupGraph()
	expected := []string{"milsons point", "wynyard", "circular quay", "st james"}

	path, err := PathSearch(graph, "milsons point", "st james")

	if err != nil {
		t.Error(err)
	}

	if fmt.Sprint(path) != fmt.Sprint(expected) {
		t.Fatalf("Expected %#v but got %#v", expected, path)
	}
}

func TestImpossibleRoute(t *testing.T) {
	graph := setupGraph()
	graph.AddNodeWithNameAndId("nowhere")

	_, err := PathSearch(graph, "circular quay", "nowhere")

	if err == nil {
		t.Fatal("Expected error but got nil")
	}
}

func TestPathSearchWithNonExistantStart(t *testing.T) {
	graph := setupGraph()

	_, err := PathSearch(graph, "north sydney", "central")

	if err == nil {
		t.Fatal("Expected error but got nil")
	}
}

func TestPathSearchWithNonExistantEnd(t *testing.T) {
	graph := setupGraph()

	_, err := PathSearch(graph, "circular quay", "north sydney")

	if err == nil {
		t.Fatal("Expected error but got nil")
	}
}

func BenchmarkPathSearch(b *testing.B) {
	graph := setupGraph()

	for n := 0; n < b.N; n++ {
		PathSearch(graph, "milsons point", "st james")
	}
}
