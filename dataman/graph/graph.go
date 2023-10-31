package graph

import (
	"container/list"
)

type Graph struct {
	Vertices int
	AdjList  []*list.List
}

func (g *Graph) AddEdge(src, dest int) {
	g.AdjList[src].PushBack(dest)
}
