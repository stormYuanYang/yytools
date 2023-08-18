// Package _const.

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
// 创建日期:2023/6/7
package base

const (
	TEN          = 10   // 十
	ONE_HUNDRED  = 100  // 百
	ONE_THOUSAND = 1000 // 千
	
	TEN_THOUSAND     = 1e4 // 万
	HUNDRED_THOUSAND = 1e5 // 十万
	MILLION          = 1e6 // 百万
	TEN_MILLION      = 1e7 // 千万
	HUNDRED_MILLION  = 1e8 // 亿
	
	// 接下来的单位可能就超过了int32的表示范围(最大的int32大约是二十一亿)
	BILLION         = 1e9  // 十亿
	TEN_BILLION     = 1e10 // 百亿
	HUNDRED_BILLION = 1e11 // 千亿
	TRILLION        = 1e12 // 万亿
)