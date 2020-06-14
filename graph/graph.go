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
	if g.Type == Directed {
		if val1, ok := g.edges[u]; ok {
			if val2, ok := val1[v]; ok {
				return val2, true
			}
			//else {
			//	//return nil, false
			//}
		}
	} else if g.Type == Undirected {
		if val1, ok := g.edges[u]; ok {
			if val2, ok := val1[v]; ok {
				return val2, true
			}
		}
		if val1, ok := g.edges[v]; ok {
			if val2, ok := val1[u]; ok {
				return val2, true
			}
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
	if g.nodes[u] == nil {
		panic("node does not exist")
	}
	for i, _ := range g.nodes {
		if data, ok := g.Edge(u, i); ok {
			f(g.nodes[i], data)
		}
		//if i != u {
		//	if data, ok := g.Edge(u, i); ok {
		//		f(g.nodes[i], data)
		//	}
		//}
	}
}

//type Int struct {
//	id int
//}
//func (d Int) ID() int { return d.id }
//
//func main() {
//	G := New(Directed)
//	G.AddNode(Int{1})
//	G.AddNode(Int{2})
//	G.AddNode(Int{3})
//	G.AddNode(Int{4})
//	G.AddNode(Int{5})
//	G.AddEdge(1, 2, nil)
//	G.AddEdge(1, 3, nil)
//	G.AddEdge(3, 4, nil)
//	G.AddEdge(5, 4, nil)
//	G.AddEdge(4, 4, nil)
//
//	G.Nodes(func(v Node){
//		G.Neighbours(v.ID(), func(k Node, edgeData interface{}) {
//			fmt.Println(v.ID())
//			fmt.Println(k, edgeData)
//			//nodes = append(nodes, v.ID())
//		})
//	})
//
////G.Edges(func(start, end Node, e interface{}) {
////	fmt.Println(start, end, e)
////})
////	fmt.Println(G.edges)
//	//fmt.Println(G.Edge(4, 4))
//	////nodes := make([]int, 0)
//	//G.Neighbours(1, func(v Node, edgeData interface{}) {
//	//	fmt.Println(v, edgeData)
//	//	//nodes = append(nodes, v.ID())
//	//})
////	fmt.Println(nodes)
//
//
//	//fmt.Println(G.edges[4])
//}
