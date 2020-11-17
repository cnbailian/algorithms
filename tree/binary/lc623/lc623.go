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

package lc623

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

// 623. 在二叉树中增加一行
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}
	if d == 1 {
		return &TreeNode{Left: root, Val: v}
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		d--
		l := len(queue)
		for i := 0; i < l; i++ {
			if d == 1 {
				queue[i].Left = &TreeNode{
					Val:   v,
					Left:  queue[i].Left,
					Right: nil,
				}
				queue[i].Right = &TreeNode{
					Val:   v,
					Left:  nil,
					Right: queue[i].Right,
				}
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return root
}
