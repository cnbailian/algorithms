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

package lc513

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

// 513. 找树左下角的值
// 首先，层序遍历最直观，每次更新最后一层的第一个值，迭代结束后直接返回
// 递归的想法应该是记录深度，结束后找到最深的第一个节点就可以
// 注意: 您可以假设树（即给定的根节点）不为 NULL。省略 root == nil 判断
func findBottomLeftValue(root *TreeNode) int {
	var res int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		res = queue[0].Val
		for i := 0; i < l; i++ {
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

func findBottomLeftValue2(root *TreeNode) int {
	dep, res := 0, 0
	var traverse func(root *TreeNode, depth int)
	traverse = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth > dep {
			dep = depth
			res = root.Val
		}
		traverse(root.Left, depth+1)
		traverse(root.Right, depth+1)
	}
	traverse(root, 1)
	return res
}
