package transportdata

// Vertex represents a graph vertex with zero or more edges.
type Vertex struct {
	VertexID string
	Edges    []*Edge
}

func newVertex(vertexID string) *Vertex {
	var edges []*Edge
	return &Vertex{vertexID, edges}
}
