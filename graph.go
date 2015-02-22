package transportdata

import "math"

type graph struct {
	nodes map[string]node
	edges map[string][]edge
}

type node struct {
	Id   string
	Name string
}

type edge struct {
	FromNode *node
	ToNode   *node
}

const maxDistance int = math.MaxInt64

func NewGraph() graph {
	nodes := make(map[string]node)
	edges := make(map[string][]edge)
	return graph{nodes, edges}
}

func (graph graph) AddNode(id, name string) {
	node := node{id, name}
	edges := make([]edge, 0)
	graph.nodes[node.Id] = node
	graph.edges[node.Id] = edges
}

func (graph graph) AddEdge(fromId, toId string) {
	from := graph.nodes[fromId]
	to := graph.nodes[toId]
	edge := edge{&from, &to}
	graph.edges[fromId] = append(graph.edges[fromId], edge)
}

func PathSearch(graph graph, startId, endId string) []string {
	distanceMap := make(map[string]int)
	unvisited := make(StringSet)
	footprints := make(map[string]string)

	for nodeId := range graph.nodes {
		distanceMap[nodeId] = maxDistance
		unvisited.Add(nodeId)
	}

	distanceMap[startId] = 0

	current := startId

	for {
		for _, edge := range graph.edges[current] {
			newDistance := distanceMap[current] + 1
			if newDistance < distanceMap[edge.ToNode.Id] {
				distanceMap[edge.ToNode.Id] = newDistance
				footprints[edge.ToNode.Id] = current
			}
		}

		unvisited.Remove(current)

		if current == endId {
			break
		}

		if len(unvisited) == 0 {
			break
		}

		// Find the next current node
		min := maxDistance
		nextNodeId := unvisited.FirstValue()

		for nodeId := range unvisited {
			distance := distanceMap[nodeId]

			if distance < min {
				min = distance
				nextNodeId = nodeId
			}
		}

		current = nextNodeId
	}

	path := []string{endId}
	key := current

	for {
		previous, ok := footprints[key]

		if ok {
			path = append(path, previous)
			key = previous
		} else {
			break
		}
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
