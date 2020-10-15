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

func TestDFSAdjacencyTable(t *testing.T) {
	var validate = []struct {
		start    string
		end      string
		expected bool
	}{
		{"you", "thom", true},
		{"you", "anuj", true},
		{"bob", "claire", false},
		{"you", "Nonexistent", false},
	}
	for _, v := range validate {
		result := DFSAdjacencyTable(at, v.start, v.end)
		if result != v.expected {
			t.Errorf("start: %s, end: %s, result: %t, expected: %t", v.start, v.end, result, v.expected)
		}
	}
}

// 不是随机图，所以测试数据仅供参考
func BenchmarkDFSAdjacencyTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DFSAdjacencyTable(at, "you", "thom")
	}
}
