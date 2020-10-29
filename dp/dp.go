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

package dp

var items = []struct {
	name   string
	weight int
	value  int
}{
	{"water", 3, 10},
	{"book", 1, 3},
	{"food", 2, 9},
	{"jacket", 2, 5},
	{"camera", 1, 6},
}
var backpackCapacity = 6

func Backpack() {
	cell := make([][]int, len(items))
	for i, item := range items {
		cell[i] = make([]int, backpackCapacity)
		for j := 1; j <= backpackCapacity; j++ {
			index := i - 1

			cell[i][index] = max(cell[i-1][j], cell[i-1][j-item.weight])
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
