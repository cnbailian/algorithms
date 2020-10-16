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

import "strconv"

// 邻接表
type AdjacencyTableUnweightedGraph map[string][]string

func (at AdjacencyTableUnweightedGraph) Graph() {
	ag := NewAsciiGraph("Adjacency Table Graph")
	var vertexes []string
	for vertex := range at {
		vertexes = append(vertexes, vertex)
	}
	ag.AddCollectionRow("vertex", vertexes)
	edges := make([][]string, len(vertexes))
	for i, vertex := range vertexes {
		adjacentVertexes := at[vertex]
		edges[i] = make([]string, len(vertexes))
		for j, v := range vertexes {
			edges[i][j] = "0"
			for _, av := range adjacentVertexes {
				if v == av {
					edges[i][j] = "1"
				}
			}
		}
	}
	ag.AddTableRows("edge", edges)
	ag.Print()
}

func (at AdjacencyTableUnweightedGraph) GetVertexDegree(vertex string) int {
	vertexes, ok := at[vertex]
	if !ok {
		return 0
	}
	return len(vertexes)
}

// 有向边
func (at AdjacencyTableUnweightedGraph) GetEdge(v, j string) bool {
	vertexes, ok := at[v]
	if !ok {
		return false
	}
	for _, vertex := range vertexes {
		if vertex == j {
			return true
		}
	}
	return false
}

type AdjacencyTableWeightedGraph map[string]map[string]int

func (atw AdjacencyTableWeightedGraph) Graph() {
	ag := NewAsciiGraph("Adjacency Table Weighted Graph")
	var vertexes []string
	for vertex := range atw {
		vertexes = append(vertexes, vertex)
	}
	ag.AddCollectionRow("vertex", vertexes)
	edges := make([][]string, len(vertexes))
	for i, vertex := range vertexes {
		adjacentVertexes := atw[vertex]
		edges[i] = make([]string, len(vertexes))
		for j, v := range vertexes {
			edges[i][j] = "0"
			for av, weight := range adjacentVertexes {
				if v == av {
					edges[i][j] = strconv.Itoa(weight)
				}
			}
		}
	}
	ag.AddTableRows("edge", edges)
	ag.Print()
}

func (atw AdjacencyTableWeightedGraph) GetVertexDegree(vertex string) int {
	vertexes, ok := atw[vertex]
	if !ok {
		return 0
	}
	return len(vertexes)
}

// 0 为不可达
func (atw AdjacencyTableWeightedGraph) GetEdgeWeight(v, j string) int {
	vertexes, ok := atw[v]
	if !ok {
		return 0
	}
	for vertex, weight := range vertexes {
		if vertex == j {
			return weight
		}
	}
	return 0
}
