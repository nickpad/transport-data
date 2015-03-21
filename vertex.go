package transportdata

type vertex struct {
	vertexID string
	edges    []*edge
}

func newVertex(vertexID string) *vertex {
	var edges []*edge
	return &vertex{vertexID, edges}
}
