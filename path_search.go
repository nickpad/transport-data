package transportdata

import (
	"fmt"
	"math"
)

const maxDistance int = math.MaxInt64

type pathSearchState struct {
	distances   map[string]int
	breadcrumbs map[string]string
	unvisited   StringSet
	current     string
	destination string
}

func newPathSearchState(graph graph, startId, endId string) *pathSearchState {
	distances := make(map[string]int)
	footprints := make(map[string]string)
	unvisited := make(StringSet)
	return &pathSearchState{distances, footprints, unvisited, startId, endId}
}

func (state *pathSearchState) prepare(graph graph) {
	for nodeId := range graph.nodes {
		state.distances[nodeId] = maxDistance
		state.unvisited.Add(nodeId)
	}

	state.distances[state.current] = 0
}

func (state *pathSearchState) getDistance(nodeId string) (int, error) {
	distance, ok := state.distances[nodeId]

	if ok {
		return distance, nil
	} else {
		return 0, fmt.Errorf("no distance value for node %#v", nodeId)
	}
}

func (state *pathSearchState) calculateDistances(edges []edge) error {
	for _, edge := range edges {
		currentDistance, err := state.getDistance(state.current)

		if err != nil {
			return err
		}

		newDistance := currentDistance + 1

		existingDistance, err := state.getDistance(edge.ToNode.Id)

		if err != nil {
			return err
		}

		if newDistance < existingDistance {
			state.distances[edge.ToNode.Id] = newDistance
			state.breadcrumbs[edge.ToNode.Id] = state.current
		}
	}

	state.unvisited.Remove(state.current)

	return nil
}

func (state *pathSearchState) findNext() (string, error) {
	min := maxDistance
	nextNodeId := state.unvisited.FirstValue()

	for nodeId := range state.unvisited {
		distance, err := state.getDistance(nodeId)

		if err != nil {
			return "", err
		}

		if distance < min {
			min = distance
			nextNodeId = nodeId
		}
	}

	return nextNodeId, nil
}

func (state *pathSearchState) searchIsDone() bool {
	if state.current == state.destination {
		return true
	}

	if len(state.unvisited) == 0 {
		return true
	}

	return false
}

func (state *pathSearchState) pathTo(nodeId string) []string {
	path := []string{nodeId}

	for {
		previous, ok := state.breadcrumbs[nodeId]

		if ok {
			path = append([]string{previous}, path...)
			nodeId = previous
		} else {
			break
		}
	}

	return path
}

func PathSearch(graph graph, startId, endId string) ([]string, error) {
	state := newPathSearchState(graph, startId, endId)
	state.prepare(graph)

	for {
		edges, ok := graph.edges[state.current]

		if !ok {
			return nil, fmt.Errorf("no edges found for node %#v", state.current)
		}

		err := state.calculateDistances(edges)

		if err != nil {
			return nil, err
		}

		if state.searchIsDone() {
			break
		}

		nextId, err := state.findNext()

		if err != nil {
			return nil, err
		}

		state.current = nextId
	}

	return state.pathTo(endId), nil
}
