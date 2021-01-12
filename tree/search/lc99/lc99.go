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

package lc99

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

// 99. 恢复二叉搜索树
// 中序遍历时是一个有序列表，所以找出降序对就能知道错误节点。
// 只有一个降序对时说明两个错误节点相联，有两个降序对时说明错误节点被分开。
// 这种解法不符合进阶要求，因为递归栈的存在，空间复杂度是 O(N)
func recoverTree(root *TreeNode) {
	var pre *TreeNode
	var pair []*TreeNode
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		if pre != nil && root.Val < pre.Val {
			// 此处为了表明降序对数量，将所有降序对都放入临时存储
			// 更好的做法是判断是否有第一个降序对，如果有就进行替换
			pair = append(pair, pre, root)
		}
		pre = root
		traverse(root.Right)
	}
	traverse(root)
	if len(pair) == 2 {
		pair[0].Val, pair[1].Val = pair[1].Val, pair[0].Val
	} else {
		pair[0].Val, pair[3].Val = pair[3].Val, pair[0].Val
	}
}
