package transportdata

import "container/heap"

const maxDistance int = 1000000

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

type vertex struct {
	vertexId    string
	distance    int
	successors  []*vertex
	predecessor *vertex
}

func (vtx *vertex) getWeight(destination *vertex) int {
	return 1
}

func newVertex(vertexId string) *vertex {
	successors := make([]*vertex, 0)
	return &vertex{vertexId, maxDistance, successors, nil}
}

func preparePriotityQueue(graph graph, start *vertex) (*PriorityQueue, map[string]*Item) {
	start.distance = 0

	pq := make(PriorityQueue, len(graph))
	itemMap := make(map[string]*Item)

	i := 0
	for _, vtx := range graph {
		pqItem := &Item{
			value:    vtx.vertexId,
			priority: maxDistance - vtx.distance,
			index:    i,
		}
		pq[i] = pqItem
		itemMap[vtx.vertexId] = pqItem
		i++
	}

	heap.Init(&pq)

	return &pq, itemMap
}

func shortestPathSearch(graph graph, start *vertex) *PriorityQueue {
	pq, itemMap := preparePriotityQueue(graph, start)

	for pq.Len() > 0 {
		pqItem := heap.Pop(pq).(*Item)
		currentVert := graph[pqItem.value]

		for _, successor := range currentVert.successors {
			newDistance := currentVert.distance + currentVert.getWeight(successor)
			if newDistance < successor.distance {
				successor.distance = newDistance
				successor.predecessor = currentVert
				item, ok := itemMap[successor.vertexId]
				if ok {
					pq.IncreasePriority(item, newDistance)
				}
			}
		}
	}
	return pq
}

func (vtx *vertex) pathFromStart() []vertex {
	path := []vertex{}
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
