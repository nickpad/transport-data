package transportdata

type graph map[string]*Vertex

func (graph graph) addVertex(vertexID string) *Vertex {
	vtx := newVertex(vertexID)
	graph[vertexID] = vtx
	return vtx
}

func (graph graph) connectVertices(from, to string) error {
	fromVertex, ok := graph[from]

	if !ok {
		fromVertex = graph.addVertex(from)
	}

	toVertex, ok := graph[to]

	if !ok {
		toVertex = graph.addVertex(to)
	}

	fromEdge := Edge{fromVertex, toVertex, 0, 1}
	toEdge := Edge{toVertex, fromVertex, 0, 1}

	fromVertex.Edges = append(fromVertex.Edges, &fromEdge)
	toVertex.Edges = append(toVertex.Edges, &toEdge)

	return nil
}
