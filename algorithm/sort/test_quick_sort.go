// Package sort.

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
// 创建日期:2023/10/24
package sort

import (
	"fmt"
	"github.com/stormYuanYang/yytools/algorithm/math_tools/random"
	"github.com/stormYuanYang/yytools/common/assert"
)

func QuickSortTest(cnt int) {
	arr := make([]int32, 10)
	for j := 0; j < cnt; j++ {
		for i := 0; i < len(arr); i++ {
			arr[i] = random.RandInt32(1, 1e5)
		}
		QuickSort(arr)
		for z := 1; z < len(arr); z++ {
			// 判断排序结束后是否升序
			assert.Assert(arr[z-1] <= arr[z])
		}
		//for _, v := range arr {
		//	fmt.Printf("%d\t", v)
		//}
		//println()
	}
	fmt.Printf("快速排序测试完毕..\n")
}

func QuickSortDescTest(cnt int) {
	arr := make([]int32, 100)
	for j := 0; j < cnt; j++ {
		for i := 0; i < len(arr); i++ {
			arr[i] = random.RandInt32(1, 1e5)
		}
		QuickSortDesc(arr)
		for z := 1; z < len(arr); z++ {
			// 判断排序结束后是否降序序
			assert.Assert(arr[z-1] >= arr[z])
		}
		//for _, v := range arr {
		//	fmt.Printf("%d\t", v)
		//}
		//println()
	}
	fmt.Printf("快速排序(降序)测试完毕..\n")
}