// Package test.

// 版权所有(Copyright)[yangyuan]
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// 作者:  yangyuan
// 创建日期:2023/6/14
package test

import (
	"fmt"
	"yytools/tools/algorithm/math"
)

func VoseAliasMethodTest() {
	x := []int{1, 2, 5, 6, 7, 9, 50}
	total := 0
	for _, v := range x {
		total += v
	}
	println("理论概率分布：")
	for _, v := range x {
		fmt.Printf("%.4f\t", float64(v)/float64(total))
	}
	println()

	N := 1000000
	gen := math.NewVoseAliasMethod(x)
	tmp := make(map[int]int)
	for i := 0; i < N; i++ {
		index := gen.Generation()
		tmp[index]++
	}
	println("实际输出概率分布：")
	for i, _ := range x {
		fmt.Printf("%.4f\t", float64(tmp[i])/float64(N))
	}
}