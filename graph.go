package transportdata

type graph map[string]*vertex

func (graph graph) addVertex(vertexId string) *vertex {
	vtx := newVertex(vertexId)
	graph[vertexId] = vtx
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

	fromVertex.successors = append(fromVertex.successors, toVertex)
	toVertex.successors = append(toVertex.successors, fromVertex)

	return nil
}
