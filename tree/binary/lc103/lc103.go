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

package lc103

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

// 103. 二叉树的锯齿形层次遍历
// 与 102 相比就是多了层数的奇偶数判断
func zigzagLevelOrder(root *TreeNode) (res [][]int) {
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
			if len(res)%2 == 0 {
				res = append(res, s)
				continue
			}
			res = append(res, []int{})
			for i := len(s) - 1; i >= 0; i-- {
				res[len(res)-1] = append(res[len(res)-1], s[i])
			}
		}
	}
	return
}
