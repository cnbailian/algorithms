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

func QuickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	datum := data[0]
	var left, right []int
	for i := 1; i < len(data); i++ {
		if data[i] >= datum {
			right = append(right, data[i])
		} else {
			left = append(left, data[i])
		}
	}
	return append(QuickSort(left), append([]int{datum}, QuickSort(right)...)...)
}
