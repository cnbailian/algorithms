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

var gcdCases = []struct {
	dividend int
	divisor  int
	expected int
}{
	{1680, 640, 80},
}

func TestGCD(t *testing.T) {
	for _, v := range gcdCases {
		result := GCD(v.dividend, v.divisor)
		if result != v.expected {
			t.Errorf("dividend: %d, divisor: %d, result: %d, expected: %d", v.dividend, v.divisor, result, v.expected)
		}
	}
}
