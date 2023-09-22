package graph

import (
	"fmt"
	"sort"
)

// Edge represents an edge in the graph with a source, destination, and weight.
type Edge struct {
	Src, Dest, Weight int
}

// Graph represents a graph with a list of edges.
type Graph__ struct {
	Vertices int
	Edges    []Edge
}

// Kruskal finds the minimum spanning tree of the graph.
func (g *Graph__) Kruskal() []Edge {
	mst := make([]Edge, 0)
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	parent := make([]int, g.Vertices)
	for i := 0; i < g.Vertices; i++ {
		parent[i] = i
	}

	for _, edge := range g.Edges {
		rootSrc := find(parent, edge.Src)
		rootDest := find(parent, edge.Dest)
		if rootSrc != rootDest {
			mst = append(mst, edge)
			parent[rootSrc] = rootDest
		}
	}

	return mst
}

func find(parent []int, i int) int {
	for parent[i] != i {
		i = parent[i]
	}
	return i
}

func main() {
	g := &Graph__{
		Vertices: 4,
		Edges: []Edge{
			{0, 1, 10},
			{0, 2, 6},
			{0, 3, 5},
			{1, 3, 15},
			{2, 3, 4},
		},
	}

	mst := g.Kruskal()

	fmt.Println("Minimum Spanning Tree:")
	for _, edge := range mst {
		fmt.Printf("Edge %d - %d (Weight %d)\n", edge.Src, edge.Dest, edge.Weight)
	}
}
