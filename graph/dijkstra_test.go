/*
Copyright 2020 BaiLian.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package graph

import "testing"

func TestDijkstra(t *testing.T) {
	var graphs = []struct {
		graph          AdjacencyTableWeightedGraph
		expectedPath   []string
		expectedWeight int
	}{
		{atw, []string{"fin", "A", "B", "start"}, 6},
		{
			graph: AdjacencyTableWeightedGraph{
				"start": map[string]int{
					"A": 5, "B": 2,
				},
				"A": map[string]int{
					"C": 4, "D": 2,
				},
				"B": map[string]int{
					"A": 8, "D": 7,
				},
				"C": map[string]int{
					"D": 6, "fin": 3,
				},
				"D": map[string]int{
					"fin": 1,
				},
				"fin": map[string]int{},
			},
			expectedPath:   []string{"fin", "D", "A", "start"},
			expectedWeight: 8,
		},
	}
	for i, graph := range graphs {
		resultPath, resultWeight := Dijkstra(graph.graph, "start", "fin")
		if resultWeight != graph.expectedWeight || !equalPath(resultPath, graph.expectedPath) {
			t.Errorf("graph index: %d, result path: %v, result weight: %d, expected path: %v, expected weight: %d",
				i, resultPath, resultWeight, graph.expectedPath, graph.expectedWeight)
		}
	}
}

func equalPath(path, compared []string) bool {
	if len(path) != len(compared) {
		return false
	}
	for i, s := range path {
		if s != compared[i] {
			return false
		}
	}
	return true
}

var amw = &AdjacencyMatrixWeightGraph{
	Vertexes: []string{"start", "A", "B", "fin"},
	Edges: [][]float64{
		{0, 6, 2, 0},
		{0, 0, 0, 1},
		{0, 3, 0, 5},
		{0, 0, 0, 0},
	},
}

func TestDijkstraAdjacencyMatrix(t *testing.T) {
	weights := DijkstraAdjacencyMatrix(amw, 0)
	expected := []float64{0, 5, 2, 6}
	if !equalWeights(weights, expected) {
		t.Errorf("graph: %v, weights: %v, expected: %v", amw.Vertexes, weights, expected)
	}
}

func equalWeights(weights, compared []float64) bool {
	if len(weights) != len(compared) {
		return false
	}
	for i, s := range weights {
		if s != compared[i] {
			return false
		}
	}
	return true
}
