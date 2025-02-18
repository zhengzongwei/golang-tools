/*
 * Copyright (c)2025. zhengzongwei
 * golang-tools is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *         http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

package datetime

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestConversionTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		timeSec := rand.Int63n(10000)
		f := ConversionTime(timeSec)
		fmt.Printf("%v\n", f)
	}

}

func TestTimeStr2TimeStamp(t *testing.T) {
	TimeStamp := TimeStr2TimeStamp("2024-01-03 16:57:17")
	fmt.Printf("%d", TimeStamp)
}
func TestTimeStamp2TimeStr(t *testing.T) {
	TimeStr := TimeStamp2TimeStr(1704272237)
	fmt.Printf("%s", TimeStr)
}
