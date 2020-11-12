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

package lc530

import "math"

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

// 530. 二叉搜索树的最小绝对差 && 783. 二叉搜索树节点最小距离
// 限定条件，非负值的二叉搜索树
// 二叉搜索树可以顺序排队，那么 nodes[n] - nodes[n-1] 的绝对值一定比 nodes[n] - nodes[:n-2]... 的值要小，所以时间上可以 O(n)
func getMinimumDifference(root *TreeNode) int {
	var res = math.MaxInt64
	var traverse func(root *TreeNode)
	var pre *TreeNode
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		if pre != nil && (root.Val-pre.Val) < res {
			res = root.Val - pre.Val
		}
		pre = root
		traverse(root.Right)
	}
	traverse(root)
	return res
}

// 递归放入数组，每次放入与前面的值对比，寻找最小绝对差
// 时间效率太慢，因为没有用到非负值二叉搜索树特性
// 但可以解决有负值的最小绝对差值问题
func getMinimumDifference2(root *TreeNode) int {
	var res int
	var vals []int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = getMinDifference(vals, root.Val, res)
		vals = append(vals, root.Val)
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return res
}

func getMinDifference(vals []int, nodeVal int, dVal int) int {
	if vals == nil {
		return math.MaxInt64
	}
	for _, val := range vals {
		dVal = absMin(nodeVal-val, dVal)
	}
	return dVal
}

func absMin(a, b int) int {
	a = int(math.Abs(float64(a)))
	if a < b {
		return a
	}
	return b
}
