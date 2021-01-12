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

package lc606

import (
	"strconv"
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

// 606. 根据二叉树创建字符串
func tree2str(t *TreeNode) string {
	res := &strings.Builder{}
	if t == nil {
		return res.String()
	}
	res.WriteString(strconv.Itoa(t.Val))
	if t.Left == nil && t.Right == nil {
		return res.String()
	}
	res.WriteString("(")
	res.WriteString(tree2str(t.Left))
	res.WriteString(")")
	if t.Right != nil {
		res.WriteString("(")
		res.WriteString(tree2str(t.Right))
		res.WriteString(")")
	}
	return res.String()
}
