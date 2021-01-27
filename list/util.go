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

package list

func SliceToList(s []int) *ListNode {
	res := &ListNode{}
	current := res
	for i, v := range s {
		current.Val = v
		if i < len(s)-1 {
			current.Next = &ListNode{}
		}
		current = current.Next
	}
	return res
}

func ListToSlice(root *ListNode) []int {
	var res []int
	for root != nil {
		res = append(res, root.Val)
		root = root.Next
	}
	return res
}
