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

package lc129

import "strconv"

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

// 129. 求根到叶子节点数字之和
// 递归拼接树枝，为 nil 时证明叶子节点，放入数组。最后计算总和
// 看了下题解，可以用 *10 来避免计算字符串相加，还可以不用数组存储，直接相加即可
func sumNumbers(root *TreeNode) int {
	var s []string
	var traverse func(root *TreeNode, i string)
	traverse = func(root *TreeNode, i string) {
		if root == nil {
			return
		}
		sum := i + strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			s = append(s, sum)
			return
		}
		traverse(root.Left, sum)
		traverse(root.Right, sum)
	}
	traverse(root, "")
	var res int
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		res += i
	}
	return res
}
