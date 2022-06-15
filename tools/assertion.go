// Package tools.

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
// 创建日期:2022/6/15
package tools

import "strings"

// 断言
// 当断言失败时，调用panic
func Assert(condition bool, strList ...string) {
	if !condition {
		switch len(strList) {
		case 0:
			panic("asertion failed.")
			return
		case 1:
			// 需要打印的字符串数组长度为1，直接打印
			panic(strList[0])
			return
		default:
			// 当有多个字符串时，使用strings.Builder组合字符串(join内部就是使用的strings.Builder)
			// 多个字符串之间用" "分隔
			resultStr := strings.Join(strList, " ")
			panic(resultStr)
			return
		}
	}
}