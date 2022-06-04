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
