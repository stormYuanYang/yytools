// Package math.

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
// 创建日期:2022/12/7
package math

import (
	"fmt"
	"strconv"
	"yytools/tools/assert"
)

/*
 求最大公约数(欧几里得算法)
 计算两个非负整数x和y的最大公约数:若y是0,则最大公约数为x;
 否则,将x除以y得到余数r,x和y的最大公约数即为y和r的最大公约数.

 ⚠️ 注意,这是一个递归函数.
 时间复杂度: ? 直觉一般是常量到对数之间
 空间复杂度: O(c)
*/
func Gcd(x, y int) int {
	assert.Assert(x >= 0, fmt.Sprintf("x must greater than or equal 0, x:%d", strconv.Itoa(x)))
	assert.Assert(y >= 0, fmt.Sprintf("y must greater than or equal 0, y:%d", strconv.Itoa(y)))
	if y == 0 {
		return x
	}
	r := x % y
	return Gcd(y, r)
}