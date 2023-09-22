package graph

import (
	"fmt"
	"math"
)

// Graph represents an adjacency matrix representation of a weighted graph.
type Graph_ struct {
	Vertices int
	Matrix   [][]int
}

// Dijkstra finds the shortest path from the source vertex to all other vertices.
func (g *Graph_) Dijkstra(src int) []int {
	dist := make([]int, g.Vertices)
	visited := make([]bool, g.Vertices)

	for i := range dist {
		dist[i] = math.MaxInt64
	}

	dist[src] = 0

	for i := 0; i < g.Vertices-1; i++ {
		u := minDistance(dist, visited)
		visited[u] = true

		for v := 0; v < g.Vertices; v++ {
			if !visited[v] && g.Matrix[u][v] != 0 && dist[u]+g.Matrix[u][v] < dist[v] {
				dist[v] = dist[u] + g.Matrix[u][v]
			}
		}
	}

	return dist
}

func minDistance(dist []int, visited []bool) int {
	min := math.MaxInt64
	minIndex := -1

	for v := 0; v < len(dist); v++ {
		if !visited[v] && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}

	return minIndex
}

func Dijkst() {
	g := &Graph_{
		Vertices: 5,
		Matrix: [][]int{
			{0, 4, 0, 0, 0},
			{4, 0, 8, 0, 0},
			{0, 8, 0, 7, 0},
			{0, 0, 7, 0, 9},
			{0, 0, 0, 9, 0},
		},
	}

	src := 0
	distances := g.Dijkstra(src)

	fmt.Println("Shortest distances from vertex", src, "to all other vertices:")
	for i, distance := range distances {
		fmt.Printf("Vertex %d: %d\n", i, distance)
	}
}
