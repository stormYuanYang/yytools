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

import (
	"math"
	"math/rand"
	"strconv"
)

// 返回闭区间[low,high]中的某一个数
func RandInt32(low, high int32) int32 {
	Assert(low >= 0, "invalid low:", strconv.Itoa(int(low)))
	Assert(high >= 0, "invalid high:", strconv.Itoa(int(high)))
	if low == high {
		return low
	}
	if low > high {
		low, high = high, low
	}
	Assert(!(low == 0 && high == math.MaxInt32), "low等于0时，high不能为最大值")
	n := high - low + 1
	return rand.Int31n(n) + low
}

// 返回闭区间[low,high]中的某一个数
func RandInt64(low, high int64) int64 {
	Assert(low >= 0, "invalid low:", strconv.Itoa(int(low)))
	Assert(high >= 0, "invalid high:", strconv.Itoa(int(high)))
	if low == high {
		return low
	}
	if low > high {
		low, high = high, low
	}
	Assert(!(low == 0 && high == math.MaxInt64), "low等于0时，high不能为最大值")
	n := high - low + 1
	return rand.Int63n(n) + low
}