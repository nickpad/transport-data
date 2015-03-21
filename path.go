package transportdata

import "strings"

type path []vertex

func (p path) String() string {
	vertexIDs := []string{}
	for _, vertex := range p {
		vertexIDs = append(vertexIDs, vertex.vertexID)
	}
	return strings.Join(vertexIDs, " -> ")
}
