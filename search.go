package transportdata

import (
	"container/heap"
	"math"
)

const maxDistance int = int(math.MaxInt32)

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

func search(graph graph, start *vertex, end *vertex) {
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
}
