package transportdata

import "container/heap"

type dGraph map[string]*dVertex

func (graph dGraph) addVertex(vertexId string) *dVertex {
	vertex := newVertex(vertexId)
	graph[vertexId] = vertex
	return vertex
}

func (graph dGraph) connectVertices(from, to string) error {
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

type dVertex struct {
	vertexId    string
	distance    int
	successors  []*dVertex
	predecessor *dVertex
}

func (vertex *dVertex) getWeight(destination *dVertex) int {
	return 1
}

func newVertex(vertexId string) *dVertex {
	successors := make([]*dVertex, 0)
	return &dVertex{vertexId, maxDistance, successors, nil}
}

func preparePriotityQueue(graph dGraph, start *dVertex) (*PriorityQueue, map[string]*Item) {
	start.distance = 0

	pq := make(PriorityQueue, len(graph))
	itemMap := make(map[string]*Item)

	i := 0
	for _, vertex := range graph {
		pqItem := &Item{
			value:    vertex.vertexId,
			priority: maxDistance - vertex.distance,
			index:    i,
		}
		pq[i] = pqItem
		itemMap[vertex.vertexId] = pqItem
		i++
	}

	heap.Init(&pq)

	return &pq, itemMap
}

func djikstra(graph dGraph, start *dVertex) *PriorityQueue {
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
