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

package lc290

import "strings"

// 290. 单词规律
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	m1 := map[string]byte{}
	m2 := map[byte]string{}
	for i, word := range words {
		if p, ok := m1[word]; ok {
			if p != pattern[i] {
				return false
			}
			continue
		}
		if w, ok := m2[pattern[i]]; ok {
			if w != word {
				return false
			}
		}
		m1[word] = pattern[i]
		m2[pattern[i]] = word
	}
	return true
}
