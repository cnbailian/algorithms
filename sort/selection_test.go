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

package sort

import (
	"math/rand"
	"testing"
)

var cases = struct {
	origin   []int
	expected []int
}{
	[]int{94, 99, 64, 23, 89, 1, 10, 20},
	[]int{1, 10, 20, 23, 64, 89, 94, 99},
}

func TestSelectionSort(t *testing.T) {
	data := selectionSort(cases.origin)
	if len(data) != len(cases.expected) {
		t.Errorf("data len: %d, validate len: %d", len(data), len(cases.expected))
	}
	for i, v := range data {
		if v != cases.expected[i] {
			t.Errorf("index: %d, value: %d, expected: %d", i, v, cases.expected[i])
		}
	}
}

func TestSelectionSortSimple(t *testing.T) {
	data := selectionSortSimple(cases.origin)
	if len(data) != len(cases.expected) {
		t.Errorf("data len: %d, validate len: %d", len(data), len(cases.expected))
	}
	for i, v := range data {
		if v != cases.expected[i] {
			t.Errorf("index: %d, value: %d, expected: %d", i, v, cases.expected[i])
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	data := generateSliceData()
	for i := 0; i < b.N; i++ {
		cpy := make([]int, len(data))
		copy(cpy, data)
		selectionSortSimple(cpy)
	}
}

func generateSliceData() []int {
	var data []int
	for i := 0; i < 10000; i++ {
		data = append(data, rand.Int())
	}
	return data
}
