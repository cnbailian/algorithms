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

package lc107

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

// 107. 二叉树的层次遍历 II
// 与 102 相比就多了反转
func levelOrderBottom(root *TreeNode) (res [][]int) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var s []int
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[0] != nil {
				s = append(s, queue[0].Val)
				queue = append(queue, queue[0].Left, queue[0].Right)
			}
			queue = queue[1:]
		}
		if len(s) > 0 {
			res = append(res, s)
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return
}
