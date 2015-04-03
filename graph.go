package transportdata

type graph map[string]*Vertex

func (g graph) addVertex(vertexID string) *Vertex {
	_, ok := g[vertexID]

	if !ok {
		g[vertexID] = newVertex(vertexID)
	}

	return g[vertexID]
}
