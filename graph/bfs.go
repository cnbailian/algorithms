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

import "sync"

type queue struct {
	list []string
	lock sync.Mutex
}

func newQueue() *queue {
	return &queue{
		lock: sync.Mutex{},
	}
}

func (q *queue) Push(el string) {
	q.lock.Lock()
	q.list = append(q.list, el)
	q.lock.Unlock()
}

func (q *queue) Pop() string {
	if len(q.list) == 0 {
		return ""
	}
	q.lock.Lock()
	el := q.list[0]
	q.list = q.list[1:]
	q.lock.Unlock()
	return el
}

// 只处理有向无环图（DAG）
func BFSAdjacencyTable(graph AdjacencyTableUnweightedGraph, root string, end string) bool {
	queue := newQueue()
	if _, ok := graph[root]; !ok {
		return false
	}
	if _, ok := graph[end]; !ok {
		return false
	}
	for _, vertex := range graph[root] {
		queue.Push(vertex)
	}
	for i := 0; i < len(graph); i++ {
		vertex := queue.Pop()
		if vertex == end {
			return true
		}
		for _, v := range graph[vertex] {
			queue.Push(v)
		}
	}
	return false
}
