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

package lc652

import "fmt"

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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	treeMap := map[string]struct {
		Replica  int
		TreeNode *TreeNode
	}{}
	var postorder func(root *TreeNode) string
	postorder = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		left, right := postorder(root.Left), postorder(root.Right)
		res := serialize(left, right, root.Val)
		// 得到序列化后的树，来判断是否重复
		if _, ok := treeMap[res]; ok {
			treeMap[res] = struct {
				Replica  int
				TreeNode *TreeNode
			}{Replica: treeMap[res].Replica + 1, TreeNode: root}
		} else {
			treeMap[res] = struct {
				Replica  int
				TreeNode *TreeNode
			}{Replica: 1, TreeNode: root}
		}
		return res
	}
	postorder(root)
	// 本题没有要求计算出现次数，所以 map 存储结构可以修改
	var result []*TreeNode
	for _, node := range treeMap {
		if node.Replica > 1 {
			result = append(result, node.TreeNode)
		}
	}
	return result
}

// 按照后序格式序列化
// 如果有空指针信息，可以根据序列化结果进行反序列化
func serialize(left, right string, val int) string {
	return fmt.Sprintf("%s, %s, %d", left, right, val)
}
