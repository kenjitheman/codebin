package graph

import (
	"container/list"
	"fmt"
)

// BFS explores the graph using Breadth-First Search starting from the given vertex.
func (g *Graph) BFS(v int) {
	visited := make([]bool, g.Vertices)
	queue := list.New()

	visited[v] = true
	queue.PushBack(v)

	for queue.Len() != 0 {
		vertex := queue.Front().Value.(int)
		fmt.Printf("%d ", vertex)
		queue.Remove(queue.Front())

		for e := g.AdjList[vertex].Front(); e != nil; e = e.Next() {
			adjVertex := e.Value.(int)
			if !visited[adjVertex] {
				visited[adjVertex] = true
				queue.PushBack(adjVertex)
			}
		}
	}
}

func Tryman() {
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

	fmt.Println("BFS traversal starting from vertex 2:")
	g.BFS(2)
}
