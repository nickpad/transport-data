package transportdata

type Graph struct {
	nodes graphNodes
	edges graphEdges
}

type Stop struct {
	Id   string
	Name string
}

type distanceMap map[string]int
type graphEdges map[string]StringSet
type graphNodes map[string]Stop
type unixTime int64

const maxDistance int = 1000

func NewGraph() Graph {
	nodes := make(graphNodes)
	edges := make(graphEdges)
	return Graph{nodes, edges}
}

func (graph *Graph) AddStop(id string, name string) {
	stop := Stop{id, name}
	graph.nodes[stop.Id] = stop
}

func (graph *Graph) AddEdge(fromId string, toId string) {
	graph.initEdgesFrom(fromId)
	graph.edges[fromId].Add(toId)
}

func (graph *Graph) initEdgesFrom(stopId string) {
	if _, ok := graph.edges[stopId]; !ok {
		graph.edges[stopId] = map[string]struct{}{}
	}
}

func (graph *Graph) PathSearch(fromId, toId string) int {
	distances := make(distanceMap)
	unvisited := make(StringSet)

	for key := range graph.nodes {
		distances[key] = maxDistance
		unvisited.Add(key)
	}

	distances[fromId] = 0

	current := fromId

	for {
		for neigbour := range graph.edges[current] {
			newDistance := distances[current] + 1
			if newDistance < distances[neigbour] {
				distances[neigbour] = newDistance
			}
		}

		unvisited.Remove(current)

		if current == toId {
			break
		}

		if len(unvisited) == 0 {
			break
		}

		current = findNextStop(unvisited, distances)
	}

	return distances[toId]
}

func findNextStop(unvisited StringSet, distances distanceMap) string {
	min := maxDistance
	nextStopId := unvisited.FirstValue()

	for stopId := range unvisited {
		distance := distances[stopId]

		if distance < min {
			min = distance
			nextStopId = stopId
		}
	}

	return nextStopId
}
