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

package lc429

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

// 429. N叉树的层序遍历
func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		var s []int
		l := len(queue)
		for i := 0; i < l; i++ {
			s = append(s, queue[0].Val)
			for _, node := range queue[0].Children {
				queue = append(queue, node)
			}
			queue = queue[1:]
		}
		res = append(res, s)
	}
	return res
}
