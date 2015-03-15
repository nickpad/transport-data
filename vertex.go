package transportdata

type vertex struct {
	vertexId    string
	distance    int
	successors  []*vertex
	predecessor *vertex
}

func newVertex(vertexId string) *vertex {
	successors := make([]*vertex, 0)
	return &vertex{vertexId, maxDistance, successors, nil}
}

func (vtx *vertex) getWeight(destination *vertex) int {
	return 1
}

func (vtx *vertex) pathFromStart() path {
	path := path{}
	current := vtx

	if current.predecessor != nil {
		path = append(path, *current)
	}

	for {
		predecessor := current.predecessor

		if predecessor != nil {
			path = append([]vertex{*predecessor}, path...)
			current = predecessor
		} else {
			break
		}
	}

	return path
}
