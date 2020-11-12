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

package lc559

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

// 559. N叉树的最大深度
// 还是没学会深度问题
func maxDepth(root *Node) int {
	depth := 0
	if root == nil {
		return depth
	}
	for _, child := range root.Children {
		depth = max(maxDepth(child), depth)
	}
	return depth + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
