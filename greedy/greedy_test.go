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

package greedy

import (
	"math/rand"
	"testing"
)

func TestBroadcast(t *testing.T) {
	if result := Broadcast(map[string]set{
		"Kone":   {"id": true, "nv": true, "ut": true},
		"Ktwo":   {"wa": true, "id": true, "mt": true},
		"Kthree": {"or": true, "nv": true, "ca": true},
		"Kfour":  {"ca": true, "az": true, "la": true},
		"Kfive":  {"id": true, "nv": true, "ut": true, "mt": true, "la": true},
	}); len(result) != 4 {
		t.Errorf("result: %v, expected: %d", result, 4)
	}
}

func TestMinCost(t *testing.T) {
	str, costs := cases()
	if result := minCost(str, costs); result != 3 {
		t.Errorf("result: %d, expected: %d", result, 3)
	}
}

func cases() (string, []int) {
	var resString []rune
	var resCosts []int
	for i := 0; i < 350000; i++ {
		resString = append(resString, rune(97))
		resCosts = append(resCosts, rand.Intn(50))
	}
	return string(resString), resCosts
}
