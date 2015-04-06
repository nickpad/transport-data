package routefinder

// Edge joins two vertices and has a departure and arrival time.
type Edge struct {
	From    *Vertex
	To      *Vertex
	Departs int64
	Arrives int64
}

func (edge *Edge) Weight() int64 {
	return edge.Arrives - edge.Departs
}
