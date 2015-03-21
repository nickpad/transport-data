package transportdata

type graph map[string]*vertex

func (graph graph) addVertex(vertexID string) *vertex {
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

	fromEdge := edge{fromVertex, toVertex, 0, 1}
	toEdge := edge{toVertex, fromVertex, 0, 1}

	fromVertex.edges = append(fromVertex.edges, &fromEdge)
	toVertex.edges = append(toVertex.edges, &toEdge)

	return nil
}
