package routefinder

import "strings"

// Path is a series of edges joining one vertex to another.
type Path []Edge

func (p Path) String() string {
	vertexIDs := []string{}

	if len(p) > 0 {
		vertexIDs = append(vertexIDs, p[0].From.VertexID)
	}

	for _, edge := range p {
		vertexIDs = append(vertexIDs, edge.To.VertexID)
	}
	return strings.Join(vertexIDs, " -> ")
}
