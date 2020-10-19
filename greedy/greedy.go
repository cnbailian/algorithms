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

package greedy

type set map[string]bool

func diff(a, b set) set {
	diff := set{}
	for i := range a {
		if _, ok := b[i]; !ok {
			diff[i] = true
		}
	}
	return diff
}

func merge(a, b set) {
	for i, v := range a {
		b[i] = v
	}
}

func Broadcast(broadcasts map[string]set) (selected []string) {
	current := make(set)
	for {
		maxSetName := ""
		maxSet := make(set)
		for name, s := range broadcasts {
			length := diff(s, current)
			if len(maxSet) < len(length) {
				maxSetName = name
				maxSet = length
			}
		}
		if maxSetName == "" {
			return
		}
		merge(maxSet, current)
		delete(broadcasts, maxSetName)
		selected = append(selected, maxSetName)
	}
}

// leetcode 1221.分割平衡字符串
func balancedStringSplit(s string) int {
	R, L, res := 0, 0, 0
	for _, v := range s {
		if v == 'R' {
			R++
		} else {
			L++
		}
		if R == L {
			res++
		}
	}
	return res
}

// leetcode 1578.避免重复字母的最小删除成本
func minCost(s string, cost []int) int {
	res := 0
	repeated, repeatedMaxCost, repeatedCostCount := s[0], cost[0], 0
	var isRepeated = false
	for i := 1; i < len(s); i++ {
		if s[i] == repeated {
			if cost[i] > repeatedMaxCost {
				repeatedCostCount += repeatedMaxCost
				repeatedMaxCost = cost[i]
			} else {
				repeatedCostCount += cost[i]
			}
			isRepeated = true
			continue
		}
		if isRepeated == true {
			res += repeatedCostCount
			isRepeated = false
		}
		repeated, repeatedMaxCost, repeatedCostCount = s[i], cost[i], 0
	}
	if isRepeated == true {
		res += repeatedCostCount
	}
	return res
}
