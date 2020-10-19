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

// 力扣 734.网络延迟事件
// 所有的边 times[i] = (u, v, w) 都有 1 <= u, v <= N 且 0 <= w <= 100。
// times[i] = (u, v, w)，其中 u 是源节点，v 是目标节点， w 是一个信号从源节点传递到目标节点的时间。
// 有 N 个网络节点，标记为 1 到 N。 现在，我们从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1。
func networkDelayTime(times [][]int, N int, K int) int {
	var maxWeight = 101

	// 邻接表
	table := map[int]map[int]int{}
	for i := 1; i <= N; i++ {
		table[i] = nil
	}
	for _, time := range times {
		u, v, w := time[0], time[1], time[2]
		if table[u] == nil {
			table[u] = make(map[int]int)
		}
		table[u][v] = w
	}

	// 权重表
	costs := map[int]int{}
	for u := range table {
		// 默认不可达
		costs[u] = maxWeight
		if u == K {
			costs[u] = 0
		}
	}
	// 重复节点
	processed := []int{K}

	for K >= 0 {
		cost := costs[K]
		// 更新权重表
		for v, w := range table[K] {
			newCost := cost + w
			if costs[v] == maxWeight || costs[v] > newCost {
				costs[v] = newCost
			}
		}
		processed = append(processed, K)

		// 选择相邻节点中下一个最短节点
		lowestCost := maxWeight
		K = -1
		for v, w := range costs {
			if w < lowestCost && !inProcessedNode(v, processed) {
				lowestCost = w
				K = v
			}
		}
	}

	max := -1
	for _, w := range costs {
		if w == maxWeight {
			return -1
		}
		if w > max {
			max = w
		}
	}

	return max
}

func inProcessedNode(vertex int, processed []int) bool {
	for _, v := range processed {
		if v == vertex {
			return true
		}
	}
	return false
}

// 邻接矩阵实现
func networkDelayTimeMatrix(times [][]int, N int, K int) int {
	var maxWeight = 101

	// 邻接矩阵
	var matrix = make([][]int, N+1)
	for _, time := range times {
		u, v, w := time[0], time[1], time[2]
		if w == 0 {
			// 避免权重与默认值冲突
			w = maxWeight
		}
		if matrix[u] == nil {
			matrix[u] = make([]int, N+1)
		}
		matrix[u][v] = w
	}

	// 权重表
	costs := map[int]int{}
	for u := range matrix {
		if u == 0 {
			continue
		}
		costs[u] = maxWeight
		if u == K {
			costs[u] = 0
		}
	}
	// 重复记录
	processed := []int{K}

	for K >= 0 {
		// 更新权重表
		cost := costs[K]
		for v, w := range matrix[K] {
			if w == 0 {
				continue
			}
			if w == maxWeight {
				w = 0
			}
			newCost := cost + w
			if costs[v] == maxWeight || costs[v] > newCost {
				costs[v] = newCost
			}
		}
		processed = append(processed, K)

		// 下一个最短节点
		lowestCost := maxWeight
		K = -1
		for v, w := range costs {
			if w < lowestCost && !inProcessedNode(v, processed) {
				lowestCost = w
				K = v
			}
		}
	}

	max := -1
	for _, w := range costs {
		if w == maxWeight {
			return -1
		}
		if w > max {
			max = w
		}
	}

	return max
}
