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

package recursion

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciLoop(n int) int {
	var data = []int{0, 1}
	if n < 2 {
		return data[n]
	}

	for i := 2; i < n+1; i++ {
		data = append(data, data[i-1]+data[i-2])
	}
	return data[n]
}
