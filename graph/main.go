package main

import (
	"fmt"
	"sort"
)

// Edge represents an edge in the graph
type Edge struct {
	src, dst, weight int
}

// Graph represents a weighted undirected graph
type Graph struct {
	vertices int
	edges    []Edge
}

// DisjointSet represents a disjoint-set data structure
type DisjointSet struct {
	parent []int
	rank   []int
}

// NewGraph creates a new graph with v vertices
func NewGraph(v int) *Graph {
	return &Graph{
		vertices: v,
		edges:    []Edge{},
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(src, dst, weight int) {
	g.edges = append(g.edges, Edge{src, dst, weight})
}

// NewDisjointSet creates a new disjoint-set
func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DisjointSet{parent, rank}
}

// Find finds the representative of a set
func (ds *DisjointSet) Find(i int) int {
	if ds.parent[i] != i {
		ds.parent[i] = ds.Find(ds.parent[i])
	}
	return ds.parent[i]
}

// Union unions two sets
func (ds *DisjointSet) Union(x, y int) {
	xroot := ds.Find(x)
	yroot := ds.Find(y)

	if ds.rank[xroot] < ds.rank[yroot] {
		ds.parent[xroot] = yroot
	} else if ds.rank[xroot] > ds.rank[yroot] {
		ds.parent[yroot] = xroot
	} else {
		ds.parent[yroot] = xroot
		ds.rank[xroot]++
	}
}

// KruskalMST applies Kruskal's algorithm to find the Minimum Spanning Tree
func (g *Graph) KruskalMST() []Edge {
	result := []Edge{}
	ds := NewDisjointSet(g.vertices)

	// Sort edges by weight
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].weight < g.edges[j].weight
	})

	for _, edge := range g.edges {
		x := ds.Find(edge.src)
		y := ds.Find(edge.dst)

		if x != y {
			result = append(result, edge)
			ds.Union(x, y)
		}
	}

	return result
}

func main() {
	g := NewGraph(4)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 6)
	g.AddEdge(0, 3, 5)
	g.AddEdge(1, 3, 15)
	g.AddEdge(2, 3, 4)

	mst := g.KruskalMST()

	fmt.Println("Edges in the Minimum Spanning Tree:")
	for _, edge := range mst {
		fmt.Printf("%d -- %d == %d\n", edge.src, edge.dst, edge.weight)
	}
}
