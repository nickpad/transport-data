package transportdata

type graph struct {
	nodes map[string]node
	edges map[string][]edge
}

type node struct {
	Id   string
	Name string
}

type edge struct {
	FromNode *node
	ToNode   *node
}

func NewGraph() graph {
	nodes := make(map[string]node)
	edges := make(map[string][]edge)
	return graph{nodes, edges}
}

func (graph graph) AddNode(id, name string) {
	node := node{id, name}
	edges := make([]edge, 0)
	graph.nodes[node.Id] = node
	graph.edges[node.Id] = edges
}

func (graph graph) AddEdge(fromId, toId string) {
	from := graph.nodes[fromId]
	to := graph.nodes[toId]
	edge := edge{&from, &to}
	graph.edges[fromId] = append(graph.edges[fromId], edge)
}
