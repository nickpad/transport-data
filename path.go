package transportdata

import "strings"

type path []vertex

func (p path) String() string {
	vertexIds := []string{}
	for _, vertex := range p {
		vertexIds = append(vertexIds, vertex.vertexId)
	}
	return strings.Join(vertexIds, " -> ")
}
