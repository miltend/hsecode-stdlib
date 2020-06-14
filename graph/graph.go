package graph

//package main

type Type int
type Node interface {
	ID() int
}

type Graph struct {
	Type  Type
	nodes map[int]Node
	edges map[int]map[int]interface{}
}

const (
	Directed   Type = iota
	Undirected Type = iota
)

func New(graphType Type) *Graph {
	G := &Graph{
		Type:  graphType,
		nodes: make(map[int]Node),
		edges: make(map[int]map[int]interface{}),
	}
	return G
}

func (g *Graph) AddNode(node Node) {
	id := node.ID()
	g.nodes[id] = node
}

func (g *Graph) AddEdge(u, v int, edgeData interface{}) {
	if g.nodes[u] == nil || g.nodes[v] == nil {
		panic("node does not exist")
	}
	child, ok := g.edges[u]
	if !ok {
		child = make(map[int]interface{})
		g.edges[u] = child
	}
	child[v] = edgeData
}

func (g *Graph) Edge(u, v int) (interface{}, bool) {
	//if g.lookup[u] == nil || g.lookup[v] == nil {
	//	panic("node does not exist")
	//}
	if g.Type == Directed {
		if data := g.edges[u][v]; data != nil {
			return data, true
		} else {
			return data, false
		}
	} else if g.Type == Undirected {
		data1 := g.edges[u][v]
		data2 := g.edges[v][u]
		if data1 != nil {
			return data1, true
		} else if data2 != nil {
			return data2, true
		} else {
			return nil, false
		}
	}
	return nil, false
}

func (g *Graph) Node(id int) (Node, bool) {
	if data := g.nodes[id]; data != nil {
		return data, true
	} else {
		return data, false
	}
}

func (g *Graph) Nodes(f func(Node)) {
	for _, value := range g.nodes {
		f(value)
	}
}

func (g *Graph) Edges(f func(u, v Node, edgeData interface{})) {
	for i, _ := range g.edges {
		for k, _ := range g.edges[i] {
			f(g.nodes[i], g.nodes[k], g.edges[i][k])
		}
	}
}

func (g *Graph) Neighbours(u int, f func(v Node, edgeData interface{})) {
	for i, _ := range g.nodes {
		if i != u {
			if data, ok := g.Edge(u, i); ok {
				f(g.nodes[i], data)
			}
		}
	}
}

//type Person struct {
//	Id   int
//	Name string
//}
//
//type Int struct {
//	id int
//}
//func (d Int) ID() int { return d.id }
//func (d Person) ID() int { return d.Id }

//func main() {
//	G := New(Undirected)
//	G.AddNode(Person{
//		Id:   1,
//		Name: "Fred Weasley",
//	})
//	G.AddNode(Person{
//		Id:   2,
//		Name: "Ginny Weasley",
//	})
//	G.AddNode(Person{
//		Id:   3,
//		Name: "Persi Weasley",
//	})
//	G.AddEdge(1, 2, "brother and sister")
//	G.AddEdge(1, 3, "brothers")
//	G.AddEdge(2, 3, "brother and sist")
//G.AddNode(Int{4})
//G.AddNode(Int{5})
//G.AddEdge(4, 5, "forward edge")
//G.AddEdge(5, 4, "backward edge")
//G.Neighbours(3, func(v Node, edgeData interface{}) {
//	fmt.Println(v, edgeData)
//})

//G.Edges(func(start, end Node, e interface{}) {
//	fmt.Println(start, end, e)
//})
//}
