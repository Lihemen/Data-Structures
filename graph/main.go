package main

import "fmt"

// EXERCISE: IMPLEMENT ADJACENCY MATRIX

// METHODS OF REPRESENTING GRAPHS
/**
* 1) ADJACENCY LIST
* 2) ADJACENCY MATRIX
 */

/**
TYPES OF GRAPHS Structure
1. Weighted Graph
2. Trees (Graph with a lot of restrictions)
3. Directed Graph
4. Undirected Graph (Bidirectional Graph)
5. Cyclic Graph
*/

// Graph Structure
type Graph struct {
	vertices []*Vertex
}

// Vertex Structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// Get vertex address
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check errors
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// Add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex returns a pointer to the vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// Contain checks if a key already exists as a vertex
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
}

func main() {
	test := &Graph{}

	for i := 1; i < 6; i++ {
		test.AddVertex(i)
	}
	test.AddEdge(1, 5)
	test.AddEdge(1, 5)
	test.AddEdge(3, 2)
	test.AddEdge(2, 4)
	test.Print()
}
