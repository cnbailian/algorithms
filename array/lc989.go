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

package array

// 989. 数组形式的整数加法
func addToArrayForm(A []int, K int) []int {
	i := len(A) - 1
	for K > 0 {
		// 溢出加到首位
		if i < 0 {
			A = append([]int{0}, A...)
			i++
		}
		sum := K%10 + A[i]
		// 去除最后一位
		K /= 10
		// 计算进位
		if sum/10 > 0 {
			K += sum / 10
		}
		A[i] = sum % 10
		i--
	}
	return A
}
