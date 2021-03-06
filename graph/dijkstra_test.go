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

func TestNetworkDelayTime(t *testing.T) {
	if result := networkDelayTime([][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
	}, 4, 2); result != 2 {
		t.Errorf("result: %d, expected: %d", result, 2)
	}
}
