package transportdata

import (
	"container/heap"
	"math"
)

const maxDistance int = int(math.MaxInt32)

type state struct {
	start         *vertex
	distances     map[string]int
	predecessors  map[string]*vertex
	priorityQueue *PriorityQueue
	itemMap       map[string]*Item
}

func prepareSearch(graph graph, start *vertex) *state {
	pq := make(PriorityQueue, len(graph))

	state := state{
		start:         start,
		distances:     map[string]int{},
		predecessors:  map[string]*vertex{},
		priorityQueue: &pq,
		itemMap:       map[string]*Item{},
	}

	i := 0
	for _, vtx := range graph {
		pqItem := &Item{
			value:    vtx.vertexId,
			priority: maxDistance - state.getDistance(vtx.vertexId),
			index:    i,
		}

		pq[i] = pqItem
		state.itemMap[vtx.vertexId] = pqItem
		i++
	}

	heap.Init(&pq)

	return &state
}

func (state *state) getDistance(vertexId string) int {
	if vertexId == state.start.vertexId {
		return 0
	}

	distance, ok := state.distances[vertexId]

	if ok {
		return distance
	} else {
		return maxDistance
	}
}

func (state *state) increasePriority(vertex *vertex, amount int) {
	item, ok := state.itemMap[vertex.vertexId]
	if ok {
		state.priorityQueue.IncreasePriority(item, amount)
	}
}

func (state *state) search(graph graph, end *vertex) {
	for state.priorityQueue.Len() > 0 {
		pqItem := heap.Pop(state.priorityQueue).(*Item)
		currentVert := graph[pqItem.value]
		for _, successor := range currentVert.successors {
			currentDistance := state.getDistance(currentVert.vertexId)
			successorDistance := state.getDistance(successor.vertexId)
			newDistance := currentDistance + 1 // Using 1 as a placeholder weight.
			if newDistance < successorDistance {
				state.distances[successor.vertexId] = newDistance
				state.predecessors[successor.vertexId] = currentVert
				state.increasePriority(successor, newDistance)
			}
		}
	}
}

func (state *state) pathTo(vtx *vertex) path {
	path := path{}
	current := vtx
	predecessor := state.predecessors[current.vertexId]

	if predecessor != nil {
		path = append(path, *current)
	}

	for {
		predecessor := state.predecessors[current.vertexId]

		if predecessor != nil {
			path = append([]vertex{*predecessor}, path...)
			current = predecessor
		} else {
			break
		}
	}

	return path
}
