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
 *
 */

package command

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetCommandOutput(t *testing.T) {
	// 这是一个示例测试，实际测试可能需要根据具体环境和命令调整
	command := "echo Hello"
	expected := "Hello"
	result := getCommandOutput(command)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGenerateMachineCode(t *testing.T) {
	hardwareInfo := NewHardwareInfo()
	hardwareInfo.CollectInfo()
	hardwareInfo.DisplayInfo()
	result := hardwareInfo.GenerateMachineCode()
	fmt.Printf("Generated Machine Code: %s\n", result) // 打印生成的机器码
	//if result != fmt.Sprintf("%x", expected) {
	//	t.Errorf("expected %v, got %v", fmt.Sprintf("%x", expected), result)
	//}
}

//func TestGenerateKey(t *testing.T) {
//	machineCode := "123456"
//	expected := sha256.Sum256([]byte(machineCode))
//	hardwareInfo := HardwareInfo{}
//	result := hardwareInfo.GenerateKey(machineCode)
//	if result != fmt.Sprintf("%x", expected) {
//		t.Errorf("expected %v, got %v", fmt.Sprintf("%x", expected), result)
//	}
//}

func TestCollectInfo(t *testing.T) {
	hardwareInfo := NewHardwareInfo()
	hardwareInfo.CollectInfo()
	if diff := cmp.Diff(hardwareInfo.HardwareInfo, map[string]string{"error": "Unsupported OS"}); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func TestDisplayInfo(t *testing.T) {
	hardwareInfo := NewHardwareInfo()
	hardwareInfo.CollectInfo()
	hardwareInfo.DisplayInfo()
}
