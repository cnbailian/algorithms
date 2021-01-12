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

package lc111

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

// 111. 二叉树的最小深度
// 层次遍历最简单，也最直观
// 这里的深度判断的是根节点到叶子节点的距离
// 层次遍历可以短路，而递归总是要遍历所有节点，平均时间还是 bfs 要好
func minDepth(root *TreeNode) (res int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		res++
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[0] != nil {
				if queue[0].Left == nil && queue[0].Right == nil {
					queue = nil
					break
				}
				queue = append(queue, queue[0].Left, queue[0].Right)
			}
			queue = queue[1:]
		}
	}
	return
}
