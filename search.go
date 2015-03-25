package transportdata

import (
	"container/heap"
	"math"
)

const maxDistance int64 = math.MaxInt64

type state struct {
	graph         graph             // The graph to search
	start         *vertex           // The vertex to search from
	priorityQueue *PriorityQueue    // Used to determine next vertex to explore
	distances     map[*vertex]int64 // Current distances for each vertex
	predecessors  map[*vertex]*edge // Predecessor edge used to reach each vertex
	itemMap       map[*vertex]*Item // Maps vertices to priorityQueue Items
}

func newState(graph graph, start *vertex) *state {
	pq := make(PriorityQueue, len(graph))

	state := state{
		graph:         graph,
		start:         start,
		priorityQueue: &pq,
		distances:     map[*vertex]int64{},
		predecessors:  map[*vertex]*edge{},
		itemMap:       map[*vertex]*Item{},
	}

	i := 0
	for _, vtx := range graph {
		pqItem := &Item{
			value:    vtx.vertexID,
			priority: maxDistance - state.getDistance(vtx),
			index:    i,
		}

		pq[i] = pqItem
		state.itemMap[vtx] = pqItem
		i++
	}

	heap.Init(&pq)

	return &state
}

func (state *state) getDistance(vertex *vertex) int64 {
	if vertex == state.start {
		return 0
	}

	distance, ok := state.distances[vertex]

	if ok {
		return distance
	}

	return maxDistance
}

func (state *state) increasePriority(vertex *vertex, amount int64) {
	item, ok := state.itemMap[vertex]
	if ok {
		state.priorityQueue.IncreasePriority(item, amount)
	}
}

func (state *state) search() {
	for state.priorityQueue.Len() > 0 {
		pqItem := heap.Pop(state.priorityQueue).(*Item)
		currentVert := state.graph[pqItem.value]
		for _, edge := range currentVert.edges {
			successor := edge.to
			currentDistance := state.getDistance(currentVert)
			successorDistance := state.getDistance(successor)
			newDistance := currentDistance + edge.weight()
			if newDistance < successorDistance {
				state.distances[successor] = newDistance
				state.predecessors[successor] = edge
				state.increasePriority(successor, newDistance)
			}
		}
	}
}

func (state *state) pathTo(vtx *vertex) path {
	path := path{}
	current := vtx

	for {
		predecessor, ok := state.predecessors[current]

		if ok {
			path = append([]edge{*predecessor}, path...)
			current = predecessor.from
		} else {
			break
		}
	}

	return path
}
