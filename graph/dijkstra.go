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

import (
	"math"
)

//  Grokking Algorithms
func Dijkstra(graph AdjacencyTableWeightedGraph, start, fin string) (path []string, cost int) {
	if _, ok := graph[start]; !ok {
		return nil, 0
	}
	if _, ok := graph[fin]; !ok {
		return nil, 0
	}

	costs := map[string]int{}
	parents := map[string]string{}
	var processed []string
	for vertex := range graph {
		if vertex != start {
			costs[vertex] = math.MaxInt8
			parents[vertex] = ""
			if weight, ok := graph[start][vertex]; ok {
				costs[vertex] = weight
				parents[vertex] = start
			}
		}
	}

	node := findLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for n, v := range neighbors {
			newCost := cost + v
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode(costs, processed)
	}

	resultPath := []string{fin}
	vertex, ok := fin, true
	for ok {
		vertex, ok = parents[vertex]
		if ok {
			resultPath = append(resultPath, vertex)
		}
	}

	return resultPath, costs[fin]
}

func findLowestCostNode(costs map[string]int, processed []string) string {
	lowestCost := math.MaxInt8
	lowestCostNode := ""
	for node, cost := range costs {
		if cost < lowestCost && !inProcessed(node, processed) {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func inProcessed(value string, processed []string) bool {
	for _, v := range processed {
		if v == value {
			return true
		}
	}
	return false
}

// My
func DijkstraAdjacencyMatrix(graph *AdjacencyMatrixWeightGraph, start int) (paths []float64) {
	if graph == nil || len(graph.Vertexes) < 1 || len(graph.Vertexes) != len(graph.Edges) {
		return nil
	}

	costs := make([]float64, len(graph.Vertexes))
	for i := range graph.Vertexes {
		costs[i] = math.Inf(1)
		if i == start {
			costs[i] = 0
		}
	}
	var processed []int

	var node = start
	for node >= 0 {
		cost := costs[node]
		for vertex, weight := range graph.Edges[node] {
			if weight == 0 {
				continue
			}
			newCost := cost + weight
			if newCost < costs[vertex] {
				costs[vertex] = newCost
			}
		}
		processed = append(processed, node)

		// find lowest cost vertex
		lowestCost := math.Inf(1)
		lowestVertex := -1
		for i, v := range costs {
			if v < lowestCost && !inProcessedVertex(i, processed) {
				lowestVertex = i
			}
		}
		node = lowestVertex
	}

	return costs
}

func inProcessedVertex(vertex int, processed []int) bool {
	for _, v := range processed {
		if v == vertex {
			return true
		}
	}
	return false
}
