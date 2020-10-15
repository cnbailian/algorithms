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

var at = AdjacencyTable{
	"you":    []string{"alice", "bob", "claire"},
	"bob":    []string{"anuj", "peggy"},
	"alice":  []string{"peggy"},
	"claire": []string{"thom", "jonny"},
	"anuj":   []string{},
	"peggy":  []string{},
	"thom":   []string{},
	"jonny":  []string{},
}

func TestAdjacencyTable_Graph(t *testing.T) {
	at.Graph()
}

func TestAdjacencyTable_GetVertexDegree(t *testing.T) {
	var degrees = []struct {
		vertex   string
		expected int
	}{
		{"you", 3},
		{"bob", 2},
		{"alice", 1},
		{"claire", 2},
		{"anuj", 0},
		{"peggy", 0},
		{"thom", 0},
		{"jonny", 0},
		{"Nonexistent", 0},
	}
	for _, degree := range degrees {
		result := at.GetVertexDegree(degree.vertex)
		if result != degree.expected {
			t.Errorf("vertex: %s, result: %d, expected: %d", degree.vertex, result, degree.expected)
		}
	}
}

func TestAdjacencyTable_GetEdge(t *testing.T) {
	var edges = []struct {
		vertexes [2]string
		expected bool
	}{
		{[2]string{"you", "alice"}, true},
		{[2]string{"you", "bob"}, true},
		{[2]string{"bob", "you"}, false},
		{[2]string{"claire", "thom"}, true},
		{[2]string{"thom", "Nonexistent"}, false},
	}
	for _, edge := range edges {
		result := at.GetEdge(edge.vertexes[0], edge.vertexes[1])
		if result != edge.expected {
			t.Errorf("vertexes: %v, result: %t, expected: %t", edge.vertexes, result, edge.expected)
		}
	}
}
