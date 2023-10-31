package graph

import (
	"container/list"
	"fmt"
)

// DFS explores the graph using Depth-First Search starting from the given vertex.
func (g *Graph) DFS(v int, visited []bool) {
	visited[v] = true
	fmt.Printf("%d ", v)
	for e := g.AdjList[v].Front(); e != nil; e = e.Next() {
		adjVertex := e.Value.(int)
		if !visited[adjVertex] {
			g.DFS(adjVertex, visited)
		}
	}
}

func TryDFS() {
	g := &Graph{
		Vertices: 4,
		AdjList:  make([]*list.List, 4),
	}
	for i := range g.AdjList {
		g.AdjList[i] = list.New()
	}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	visited := make([]bool, g.Vertices)
	fmt.Println("DFS traversal starting from vertex 2:")
	g.DFS(2, visited)
}
