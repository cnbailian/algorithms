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

package lc872

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

// 872. 叶子相似的树
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var traverse func(root *TreeNode, res *[]int)
	traverse = func(root *TreeNode, res *[]int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			*res = append(*res, root.Val)
			return
		}
		traverse(root.Left, res)
		traverse(root.Right, res)
	}
	var res1, res2 []int
	traverse(root1, &res1)
	traverse(root2, &res2)
	return eq(res1, res2)
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
