package transportdata

import (
	"container/heap"
	"math"
)

// The maximum distance a vertex can be from the start vertex.
const maxDistance int64 = math.MaxInt64

// State stores the current state of a search.
type State struct {
	graph         graph             // The graph to search
	start         *Vertex           // The vertex to search from
	priorityQueue *PriorityQueue    // Used to determine next vertex to explore
	distances     map[*Vertex]int64 // Current distances for each vertex
	predecessors  map[*Vertex]*Edge // Predecessor edge used to reach each vertex
	itemMap       map[*Vertex]*Item // Maps vertices to priorityQueue Items
}

// NewState initializes a new State instance.
func NewState(graph graph, start *Vertex) *State {
	pq := make(PriorityQueue, len(graph))

	state := State{
		graph:         graph,
		start:         start,
		priorityQueue: &pq,
		distances:     map[*Vertex]int64{},
		predecessors:  map[*Vertex]*Edge{},
		itemMap:       map[*Vertex]*Item{},
	}

	i := 0
	for _, vtx := range graph {
		pqItem := &Item{
			value:    vtx.VertexID,
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

func (state *State) getDistance(vertex *Vertex) int64 {
	if vertex == state.start {
		return 0
	}

	distance, ok := state.distances[vertex]

	if ok {
		return distance
	}

	return maxDistance
}

func (state *State) increasePriority(vertex *Vertex, amount int64) {
	item, ok := state.itemMap[vertex]
	if ok {
		state.priorityQueue.IncreasePriority(item, amount)
	}
}

// Search performs a shortest-path search over the State graph.
func (state *State) Search() {
	for state.priorityQueue.Len() > 0 {
		pqItem := heap.Pop(state.priorityQueue).(*Item)
		currentVert := state.graph[pqItem.value]
		for _, edge := range currentVert.Edges {
			successor := edge.To
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

// PathTo returns the path from the State start to the given vertex.
func (state *State) PathTo(vtx *Vertex) Path {
	path := Path{}
	current := vtx

	for {
		predecessor, ok := state.predecessors[current]

		if ok {
			path = append([]Edge{*predecessor}, path...)
			current = predecessor.From
		} else {
			break
		}
	}

	return path
}
