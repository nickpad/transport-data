package routefinder

import "sort"

type (
	edgeList []*Edge

	// Vertex represents a graph vertex with zero or more edges.
	Vertex struct {
		VertexID string
		Edges    edgeList
	}
)

func newVertex(vertexID string) *Vertex {
	var edges edgeList
	return &Vertex{vertexID, edges}
}

func (vtx *Vertex) AddEdge(to *Vertex, departs, arrives int64) *Edge {
	edge := Edge{vtx, to, departs, arrives}
	vtx.Edges = append(vtx.Edges, &edge)
	return &edge
}

func (vtx *Vertex) Prepare() {
	sort.Sort(vtx.Edges)
}

func (vtx *Vertex) EdgesFrom(time int64) edgeList {
	index := sort.Search(len(vtx.Edges), func(i int) bool {
		return vtx.Edges[i].Departs >= time
	})

	return vtx.Edges[index:]
}

func (edges edgeList) Len() int {
	return len(edges)
}

func (edges edgeList) Less(i, j int) bool {
	return edges[i].Departs < edges[j].Departs
}

func (edges edgeList) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}
