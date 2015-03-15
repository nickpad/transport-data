package transportdata

import "strings"

type path []vertex

func (path *path) String() string {
	vertexIds := []string{}
	for _, vertex := range *path {
		vertexIds = append(vertexIds, vertex.vertexId)
	}
	return strings.Join(vertexIds, " -> ")
}
