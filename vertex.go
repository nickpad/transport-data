package transportdata

type vertex struct {
	vertexId   string
	successors []*vertex
}

func newVertex(vertexId string) *vertex {
	successors := make([]*vertex, 0)
	return &vertex{vertexId, successors}
}
