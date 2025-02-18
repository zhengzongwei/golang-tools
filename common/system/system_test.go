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

package system

import (
	"fmt"
	"testing"
)

func TestGetCPUinfo(t *testing.T) {
	cpuInfo := GetCPUinfo()
	fmt.Printf("%v#\n", cpuInfo)
}

func TestGetNodeInfo(t *testing.T) {
	nodeInfo := GetNodeInfo()
	fmt.Printf("%v#\n", nodeInfo)
}

func TestGetDiskinfo(t *testing.T) {
	diskInfo := GetDiskinfo()
	fmt.Printf("%v#\n", diskInfo)
}

func TestGetSysteminfo(t *testing.T) {
	systenInfo := GetSysteminfo()
	fmt.Printf("%v#\n", systenInfo.CPU)

}
