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

func findSmallest(data []int) int {
	var smallest = data[0]
	var smallestIndex = 0
	for i := 1; i < len(data); i++ {
		if data[i] < smallest {
			smallest = data[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		smallest := findSmallest(data[i:]) + i
		data[i], data[smallest] = data[smallest], data[i]
	}
	return data
}

type SelectionSort struct {
	Data []int
}

func (s *SelectionSort) Pop(index int) int {
	res := s.Data[index]
	s.Data = append(s.Data[:index], s.Data[index+1:]...)
	return res
}

func selectionSortSimple(data []int) []int {
	s := SelectionSort{Data: data}
	var result []int
	var length = len(s.Data)
	for i := 0; i < length; i++ {
		smallest := findSmallest(s.Data)
		result = append(result, s.Pop(smallest))
	}
	return result
}
