package transportdata

import (
	"container/heap"
	"math"
)

const maxDistance int64 = math.MaxInt64

type state struct {
	start         *vertex
	distances     map[string]int64
	predecessors  map[string]*vertex
	priorityQueue *PriorityQueue
	itemMap       map[string]*Item
}

func prepareSearch(graph graph, start *vertex) *state {
	pq := make(PriorityQueue, len(graph))

	state := state{
		start:         start,
		distances:     map[string]int64{},
		predecessors:  map[string]*vertex{},
		priorityQueue: &pq,
		itemMap:       map[string]*Item{},
	}

	i := 0
	for _, vtx := range graph {
		pqItem := &Item{
			value:    vtx.vertexID,
			priority: maxDistance - state.getDistance(vtx.vertexID),
			index:    i,
		}

		pq[i] = pqItem
		state.itemMap[vtx.vertexID] = pqItem
		i++
	}

	heap.Init(&pq)

	return &state
}

func (state *state) getDistance(vertexID string) int64 {
	if vertexID == state.start.vertexID {
		return 0
	}

	distance, ok := state.distances[vertexID]

	if ok {
		return distance
	}

	return maxDistance
}

func (state *state) increasePriority(vertex *vertex, amount int64) {
	item, ok := state.itemMap[vertex.vertexID]
	if ok {
		state.priorityQueue.IncreasePriority(item, amount)
	}
}

func (state *state) search(graph graph, end *vertex) {
	for state.priorityQueue.Len() > 0 {
		pqItem := heap.Pop(state.priorityQueue).(*Item)
		currentVert := graph[pqItem.value]
		for _, edge := range currentVert.edges {
			successor := edge.to
			currentDistance := state.getDistance(currentVert.vertexID)
			successorDistance := state.getDistance(successor.vertexID)
			newDistance := currentDistance + edge.weight()
			if newDistance < successorDistance {
				state.distances[successor.vertexID] = newDistance
				state.predecessors[successor.vertexID] = currentVert
				state.increasePriority(successor, newDistance)
			}
		}
	}
}

func (state *state) pathTo(vtx *vertex) path {
	path := path{}
	current := vtx
	predecessor := state.predecessors[current.vertexID]

	if predecessor != nil {
		path = append(path, *current)
	}

	for {
		predecessor := state.predecessors[current.vertexID]

		if predecessor != nil {
			path = append([]vertex{*predecessor}, path...)
			current = predecessor
		} else {
			break
		}
	}

	return path
}
