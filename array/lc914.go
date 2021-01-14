/*
Copyright 2021 BaiLian.

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

package array

// 914. 卡牌分组
func hasGroupsSizeX(deck []int) bool {
	if len(deck) == 0 {
		return false
	}
	m := map[int]int{}
	for _, i := range deck {
		if _, ok := m[i]; ok {
			m[i]++
		} else {
			m[i] = 1
		}
	}
	var first = m[deck[0]]
	for _, i := range m {
		if i <= 1 || gcd(i, first) == 1 {
			return false
		}
		first = i
	}
	return true
}

func gcd(x, y int) int {
	if y != 0 {
		return gcd(y, x%y)
	}
	return x
}
