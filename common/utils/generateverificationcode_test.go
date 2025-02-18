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

package utils

import (
	"fmt"
	"testing"
)

func TestGenerateVerificationCode(t *testing.T) {
	numberCode := GenerateVerificationCode("number", 6)
	fmt.Println(numberCode)

	letterCode := GenerateVerificationCode("letter", 6)
	fmt.Println(letterCode)

	letterNumberCode := GenerateVerificationCode("letter_number", 6)
	fmt.Println(letterNumberCode)
}
