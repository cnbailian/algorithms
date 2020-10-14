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

package recursion

import (
	"cnbailian/algorithms/sort"
	"math/rand"
	"testing"
)

var quickSortCases = []struct {
	data     []int
	expected []int
}{
	{[]int{1, 7}, []int{1, 7}},
	{[]int{33, 15, 10}, []int{10, 15, 33}},
	{[]int{33, 10, 15, 7}, []int{7, 10, 15, 33}},
	{[]int{3, 5, 2, 1, 4}, []int{1, 2, 3, 4, 5}},
	{[]int{94, 99, 64, 23, 89, 1, 10, 20}, []int{1, 10, 20, 23, 64, 89, 94, 99}},
}

func TestQuickSort(t *testing.T) {
	for _, v := range quickSortCases {
		result := QuickSort(v.data)
		if !equal(result, v.expected) {
			t.Errorf("data: %v, result: %v, expected: %v", v.data, result, v.expected)
		}
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkQuickSort(b *testing.B) {
	data := generateSliceData()
	for i := 0; i < b.N; i++ {
		cpy := make([]int, len(data))
		copy(cpy, data)
		QuickSort(cpy)
	}
}

func BenchmarkSortQuickSort(b *testing.B) {
	data := generateSliceData()
	for i := 0; i < b.N; i++ {
		cpy := make([]int, len(data))
		copy(cpy, data)
		s := sort.QuickSort{Data: cpy}
		s.Sort()
	}
}

func generateSliceData() []int {
	var data []int
	for i := 0; i < 10000; i++ {
		data = append(data, rand.Int())
	}
	return data
}
