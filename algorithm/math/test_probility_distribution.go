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
package math

import (
	"fmt"
	"time"
	"yytools/algorithm/math/random"
)

func NormalMethodTest(x []int) {
	println("普通概率生成方法:")
	
	N := 1000000
	var method IProbabilityDistribution = NewNormalMethod(x)
	tmp := make(map[int]int)
	for i := 0; i < N; i++ {
		index := method.Generate()
		tmp[index]++
	}
	println("实际输出概率分布：")
	for i, _ := range x {
		fmt.Printf("%.3f\t", float64(tmp[i])/float64(N))
	}
	println()
}

func VoseAliasMethodTest(x []int) {
	println("vose的别名方法:")
	N := 1000000
	var method IProbabilityDistribution = NewVoseAliasMethod(x)
	tmp := make(map[int]int)
	for i := 0; i < N; i++ {
		index := method.Generate()
		tmp[index]++
	}
	println("实际输出概率分布：")
	for i, _ := range x {
		fmt.Printf("%.3f\t", float64(tmp[i])/float64(N))
	}
	println()
}

var ProbabilityDistribution_handlers = []func(num []int){
	NormalMethodTest,
	VoseAliasMethodTest,
}

func ProbabilityDistributionTest(num int) {
	println("概率分布测试开始...")
	random.RandSeed(time.Now().UnixMilli())
	for i := 1; i <= num; i++ {
		fmt.Printf("第%d轮测试开始\n", i)
		handlerLength := len(ProbabilityDistribution_handlers)
		var x []int
		count := 5
		for a := 0; a < count; a++ {
			r := random.RandInt(1, 100)
			x = append(x, r)
		}
		println("理论概率分布：")
		total := 0
		for _, v := range x {
			total += v
		}
		for _, v := range x {
			fmt.Printf("%.3f\t", float64(v)/float64(total))
		}
		println()
		for k := 0; k < handlerLength; k++ {
			handler := ProbabilityDistribution_handlers[k]
			handler(x)
		}
		fmt.Printf("第%d轮测试结束\n\n", i)
	}
	println("概率分布测试完毕...")
}