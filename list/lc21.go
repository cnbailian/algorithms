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

package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 21. 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	current := res
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			current.Next = l2
			l2 = l2.Next
		} else {
			current.Next = l1
			l1 = l1.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	return res.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
	l1.Next = mergeTwoLists2(l1.Next, l2)
	return l1
}
