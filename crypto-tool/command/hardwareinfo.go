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
	"crypto/sha256"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

type HardwareInfo struct {
	OsType       string
	HardwareInfo map[string]string
}

func NewHardwareInfo() *HardwareInfo {
	hi := &HardwareInfo{
		OsType:       runtime.GOOS,
		HardwareInfo: make(map[string]string),
	}
	hi.CollectInfo()
	return hi
}

func getCommandOutput(command string) string {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(output))
}

func (hi *HardwareInfo) getLinuxInfo() {
	hi.HardwareInfo["system_serial_number"] = getCommandOutput("sudo dmidecode -s system-serial-number")
	hi.HardwareInfo["board_serial"] = getCommandOutput("sudo dmidecode -s baseboard-serial-number")
	hi.HardwareInfo["cpu_id"] = getCommandOutput("sudo dmidecode -t processor | grep 'ID' | awk '{print $2}'")
	hi.HardwareInfo["hard_disk_serial"] = getCommandOutput("sudo hdparm -I /dev/sda | grep 'Serial Number' | awk '{print $3}'")
	hi.HardwareInfo["mac_address"] = getCommandOutput("cat /sys/class/net/$(ip route show default | awk '/default/ {print $5}')/address")
}

func (hi *HardwareInfo) getWindowsInfo() {
	hi.HardwareInfo["system_serial_number"] = strings.Split(getCommandOutput("wmic bios get serialnumber"), "\n")[1]
	hi.HardwareInfo["board_serial"] = strings.Split(getCommandOutput("wmic baseboard get serialnumber"), "\n")[1]
	hi.HardwareInfo["cpu_id"] = strings.Split(getCommandOutput("wmic cpu get processorid"), "\n")[1]
	hi.HardwareInfo["hard_disk_serial"] = strings.Split(getCommandOutput("wmic diskdrive get serialnumber"), "\n")[1]

	macAddressOutput := getCommandOutput("getmac | findstr /v 'Media disconnected' | findstr /v 'Hyper-V Virtual Ethernet Adapter'")
	hi.HardwareInfo["mac_address"] = strings.Fields(strings.Split(macAddressOutput, "\n")[0])[0]
}

func (hi *HardwareInfo) getMacosInfo() {
	hi.HardwareInfo["system_serial_number"] = getCommandOutput("system_profiler SPHardwareDataType | grep 'Serial Number' | awk '{print $4}'")
	hi.HardwareInfo["board_serial"] = getCommandOutput("ioreg -l | grep IOPlatformSerialNumber | awk '{print $4}' | sed 's/\"//g'")
	hi.HardwareInfo["cpu_id"] = getCommandOutput("sysctl -n machdep.cpu.brand_string")
	hi.HardwareInfo["hard_disk_serial"] = getCommandOutput("system_profiler SPSerialATADataType | grep 'Serial Number' | awk '{print $3}'")
	hi.HardwareInfo["mac_address"] = getCommandOutput("ifconfig en0 | grep ether | awk '{print $2}'")
}

func (hi *HardwareInfo) CollectInfo() {
	switch hi.OsType {
	case "linux":
		hi.getLinuxInfo()
	case "windows":
		hi.getWindowsInfo()
	case "darwin":
		hi.getMacosInfo()
	default:
		hi.HardwareInfo["error"] = "Unsupported OS"
	}
}

func (hi *HardwareInfo) GenerateMachineCode() string {
	var infoStr string
	keys := []string{"system_serial_number", "board_serial", "cpu_id", "hard_disk_serial", "mac_address"}
	for _, key := range keys {
		infoStr += hi.HardwareInfo[key]
	}
	hash := sha256.Sum256([]byte(infoStr))
	return fmt.Sprintf("%x", hash)
}

func (hi *HardwareInfo) DisplayInfo() {
	fmt.Println("OS Type:", hi.OsType) // 打印操作系统类型
	fmt.Println("Hardware Information:")
	for key, value := range hi.HardwareInfo {
		fmt.Printf("%s: %s\n", key, value)
	}
}
