package routefinder

import (
	"container/heap"
	"math"
)

// The maximum distance a vertex can be from the start vertex.
const maxDistance int64 = math.MaxInt64

type (
	distanceMap   map[*Vertex]int64
	breadcrumbMap map[*Vertex]*Edge
	pqItemMap     map[*Vertex]*Item

	// State stores the current state of a search.
	State struct {
		graph       Graph          // The graph to search
		start       *Vertex        // The vertex to search from
		end         *Vertex        // The vertex to search to
		departAt    int64          // The departure time
		pq          *PriorityQueue // Used to determine next vertex to explore
		distances   distanceMap    // Current distances for each vertex
		breadcrumbs breadcrumbMap  // Predecessor edge used to reach each vertex
		pqItems     pqItemMap      // Maps vertices to priorityQueue Items
	}
)

// NewState initializes a new State instance.
func NewState(graph Graph, start, end *Vertex, departAt int64) *State {
	state := State{
		graph:       graph,
		start:       start,
		end:         end,
		departAt:    departAt,
		distances:   distanceMap{},
		breadcrumbs: breadcrumbMap{},
		pqItems:     pqItemMap{},
	}

	state.initPriorityQueue()

	return &state
}

// Search performs a shortest-path search over the State graph.
func (state *State) Search() {
	currentTime := state.departAt
	currentVert := state.nextVertex()

	for {
		if currentVert == nil {
			break
		}
		currentDistance := state.getDistance(currentVert)
		if currentDistance >= maxDistance {
			break
		}
		for _, edge := range currentVert.EdgesFrom(currentTime) {
			successor := edge.To
			successorDistance := state.getDistance(successor)
			newDistance := currentDistance + edge.Weight()
			if newDistance < successorDistance {
				state.distances[successor] = newDistance
				state.breadcrumbs[successor] = edge
				state.increasePriority(successor, newDistance)
				currentTime = edge.Arrives
			}
		}
		if currentVert == state.end {
			break
		}
		currentVert = state.nextVertex()
	}
}

// Path returns the path from the State start to the State end.
func (state *State) Path() Path {
	path := Path{}
	current := state.end

	for {
		predecessor, ok := state.breadcrumbs[current]

		if ok {
			path = append([]Edge{*predecessor}, path...)
			current = predecessor.From
		} else {
			break
		}
	}

	return path
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
	item, ok := state.pqItems[vertex]
	if ok {
		state.pq.IncreasePriority(item, amount)
	}
}

func (state *State) initPriorityQueue() {
	pq := make(PriorityQueue, len(state.graph))

	i := 0
	for _, vtx := range state.graph {
		pqItem := &Item{
			value:    vtx,
			priority: maxDistance - state.getDistance(vtx),
			index:    i,
		}

		pq[i] = pqItem
		state.pqItems[vtx] = pqItem
		i++
	}

	heap.Init(&pq)
	state.pq = &pq
}

func (state *State) nextVertex() *Vertex {
	if len(*state.pq) > 0 {
		item := heap.Pop(state.pq).(*Item)
		return item.value
	}
	return nil
}
