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

package lc671

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

// 671. 二叉树中第二小的节点
// 对比全部查找的优化点在于：当前节点 val > second，那么当前节点的子节点也一定大于 second，
// 所以可以忽略部分子树，那么层序遍历就比深度遍历要好
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || (root.Left == nil || root.Right == nil) {
		return -1
	}
	min := root.Val
	second := -1
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		cur := queue[0]
		if cur.Val > min && (second == -1 || cur.Val < second) {
			second = cur.Val
		}
		if cur.Left != nil && cur.Right != nil && (second == -1 || cur.Val < second) {
			queue = append(queue, cur.Left, cur.Right)
		}
		queue = queue[1:]
	}
	return second
}
