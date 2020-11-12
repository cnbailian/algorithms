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

package lc515

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

// 515. 在每个树行中找最大值
// 存在行数要求，层序遍历最直观。但递归记录深度也可以
func largestValues(root *TreeNode) []int {
	var res []int
	var traverse func(root *TreeNode, depth int)
	traverse = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) == depth {
			res = append(res, root.Val)
		} else if len(res) > depth {
			res[depth] = max(res[depth], root.Val)
		}
		traverse(root.Left, depth+1)
		traverse(root.Right, depth+1)
	}
	traverse(root, 0)
	return res
}

func largestValues2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		res = append(res, queue[0].Val)
		for i := 0; i < l; i++ {
			res[len(res)-1] = max(res[len(res)-1], queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
