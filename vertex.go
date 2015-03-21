package transportdata

type vertex struct {
	vertexID   string
	successors []*vertex
}

func newVertex(vertexID string) *vertex {
	var successors []*vertex
	return &vertex{vertexID, successors}
}
