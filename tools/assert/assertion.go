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
package assert

import (
	"fmt"
)

// 断言
// 当断言失败时，调用panic
func Assert(condition bool, list ...interface{}) {
	if !condition {
		switch len(list) {
		case 0:
			panic("asertion failed.")
			return
		default:
			str := fmt.Sprint(list)
			panic(str)
			return
		}
	}
}