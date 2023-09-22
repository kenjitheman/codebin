package graph

import (
	"container/list"
)

// Graph represents an adjacency list representation of a graph.
type Graph struct {
	Vertices int
	AdjList  []*list.List
}

// AddEdge adds an edge between two vertices.
func (g *Graph) AddEdge(src, dest int) {
	g.AdjList[src].PushBack(dest)
}
