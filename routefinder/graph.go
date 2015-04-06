package routefinder

type Graph map[string]*Vertex

func (g Graph) AddVertex(vertexID string) *Vertex {
	_, ok := g[vertexID]

	if !ok {
		g[vertexID] = newVertex(vertexID)
	}

	return g[vertexID]
}
