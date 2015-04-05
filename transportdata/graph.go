package transportdata

type Graph map[string]*Vertex

func (g Graph) addVertex(vertexID string) *Vertex {
	_, ok := g[vertexID]

	if !ok {
		g[vertexID] = newVertex(vertexID)
	}

	return g[vertexID]
}
