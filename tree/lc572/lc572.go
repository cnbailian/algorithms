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

package lc572

import (
	"fmt"
	"strings"
)

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

// 572. 另一个树的子树
// 思路：在 s 树中递归寻找每个树是否与 t 树为相同的树
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if isSameTree(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

// 100. 相同的树
func isSameTree(s *TreeNode, t *TreeNode) bool {
	// case1
	if s == nil && t == nil {
		return true
	}
	// case2 节点不统一
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

// 另一种思路，遍历得到字符串后查询字符串包含
func isSubtree2(s *TreeNode, t *TreeNode) bool {
	var traverse func(root *TreeNode, res *strings.Builder)
	traverse = func(root *TreeNode, res *strings.Builder) {
		if root == nil {
			res.WriteString("null")
			return
		}
		res.WriteString(fmt.Sprintf("(%d)", root.Val))
		traverse(root.Left, res)
		traverse(root.Right, res)
	}
	sStr, tStr := &strings.Builder{}, &strings.Builder{}
	traverse(s, sStr)
	traverse(t, tStr)
	return strings.Contains(sStr.String(), tStr.String())
}
