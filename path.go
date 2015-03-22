package transportdata

import "strings"

type path []edge

func (p path) String() string {
	vertexIDs := []string{}

	if len(p) > 0 {
		vertexIDs = append(vertexIDs, p[0].from.vertexID)
	}

	for _, edge := range p {
		vertexIDs = append(vertexIDs, edge.to.vertexID)
	}
	return strings.Join(vertexIDs, " -> ")
}
