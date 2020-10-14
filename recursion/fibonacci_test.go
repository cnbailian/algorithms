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

import "testing"

var fibonacciCases = []struct {
	number   int
	expected int
}{
	{0, 0},
	{1, 1},
	{10, 55},
	{20, 6765},
	{30, 832040},
	{40, 1.02334155e+08},
}

func TestFibonacci(t *testing.T) {
	for _, v := range fibonacciCases {
		value := Fibonacci(v.number)
		if v.expected != value {
			t.Errorf("return； %d, expected: %d", value, v.expected)
		}
	}
}

func TestFibonacciLoop(t *testing.T) {
	for _, v := range fibonacciCases {
		value := FibonacciLoop(v.number)
		if v.expected != value {
			t.Errorf("return； %d, expected: %d", value, v.expected)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(40)
	}
}

func BenchmarkFibonacciLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciLoop(40)
	}
}
