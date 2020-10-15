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

// 只处理有向无环图（DAG）
func DFSAdjacencyTable(graph AdjacencyTableUnweightedGraph, root string, end string) bool {
	if _, ok := graph[root]; !ok {
		return false
	}
	if _, ok := graph[end]; !ok {
		return false
	}
	for _, vertex := range graph[root] {
		if vertex == end {
			return true
		}
		if DFSAdjacencyTable(graph, vertex, end) {
			return true
		}
	}
	return false
}
