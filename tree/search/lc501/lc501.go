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

package lc501

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

// 501. 二叉搜索树中的众数
// 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
// 官方题解使用 Morris，还没学会
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var mode []int
	maxCount, count, pre := 0, 0, 0
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		// 发现重复数
		if root.Val == pre {
			count++
		} else {
			count = 1
		}
		// 更新 maxCount 和 mode
		if count > maxCount {
			mode = []int{root.Val}
			maxCount = count
		} else if count == maxCount {
			mode = append(mode, root.Val)
		}
		pre = root.Val
		inorder(root.Right)
	}
	inorder(root)
	return mode
}
