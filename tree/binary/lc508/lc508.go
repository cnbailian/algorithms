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

package lc508

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 508. 出现次数最多的子树元素和
func findFrequentTreeSum(root *TreeNode) []int {
	m := map[int]int{}
	// 后序遍历求出所有和与次数
	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		i := root.Val + traverse(root.Left) + traverse(root.Right)
		if _, ok := m[i]; !ok {
			m[i] = 1
		} else {
			m[i]++
		}
		return i
	}
	traverse(root)
	// 检索最高次数值
	var maxCount int
	var res []int
	for val, count := range m {
		if count > maxCount {
			maxCount = count
			res = []int{val}
		}
		if count == maxCount {
			res = append(res, val)
		}
	}
	return res
}
