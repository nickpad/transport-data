package transportdata

type edge struct {
	from    *vertex
	to      *vertex
	departs int64
	arrives int64
}

func (edge *edge) weight() int64 {
	return edge.arrives - edge.departs
}
