package graph

import (
	"fmt"
	"testing"
)

func getAdjacencyMatrixUnweightedUndirectedGraph() AdjacencyMatrixUnweightedGraph {
	am := AdjacencyMatrixUnweightedGraph{
		Vertexes: []string{"v0", "v1", "v2", "v3"},
		Edges: [][]string{
			{"0", "1", "1", "0"},
			{"1", "0", "0", "1"},
			{"1", "0", "0", "1"},
			{"0", "1", "1", "0"},
		},
	}
	return am
}

func TestAdjacencyMatrixUnweightedGraph(t *testing.T) {
	am := getAdjacencyMatrixUnweightedUndirectedGraph()
	am.Graph()
}

func TestAdjacencyMatrixUnweightedGraph_GetVertexDegree(t *testing.T) {
	am := getAdjacencyMatrixUnweightedUndirectedGraph()
	var degrees = []struct {
		vertex   int
		expected int
	}{
		{0, 2},
		{1, 2},
		{2, 2},
		{3, 2},
	}
	for _, degree := range degrees {
		if am.GetVertexDegree(degree.vertex) != degree.expected {
			t.Error(fmt.Sprintf("vertex: %s, degree: %d, expected: %d",
				am.Vertexes[degree.vertex], am.GetVertexDegree(degree.vertex), degree.expected))
		}
	}
}

func TestAdjacencyMatrixUnweightedGraph_GetEdge(t *testing.T) {
	am := getAdjacencyMatrixUnweightedUndirectedGraph()
	var edges = []struct {
		vertexes [2]int
		expected bool
	}{
		{[2]int{0, 1}, true},
		{[2]int{0, 2}, true},
		{[2]int{0, 3}, false},
		{[2]int{1, 2}, false},
		{[2]int{1, 3}, true},
		{[2]int{2, 3}, true},
	}
	for _, edge := range edges {
		if am.GetEdge(edge.vertexes[0], edge.vertexes[1]) != edge.expected {
			vertexes := am.Vertexes[edge.vertexes[0]] + "--" + am.Vertexes[edge.vertexes[1]]
			t.Error(fmt.Sprintf("vertexes: %s, expected: %t", vertexes, edge.expected))
		}
	}
}
