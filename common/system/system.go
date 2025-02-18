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

/*
 * @Title system.go
 * @Description 获取系统相关信息 CPU 网卡 系统类型等等信息
 * @Author zhengzongwei 2021/11/16 11:16
 * @Update
 */

package system

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

type DiskInfo struct {
	DiskInfo disk.PartitionStat `json:"diskInfo"`
	Usage    disk.UsageStat     `json:"usage"`
}

type NodeInfo struct {
	NODE *host.InfoStat `json:"node"`
	CPU  []cpu.InfoStat `json:"cpu"`
	DISK []DiskInfo     `json:"disk"`
}

/*
 * @title: GetCPUinfo
 * @description: 获取CPU相关信息
 * @param:
 * @return:
 */

func GetCPUinfo() []cpu.InfoStat {
	cpuInfo, _ := cpu.Info()
	return cpuInfo
	//fmt.Printf("%v#\n", cpuInfo)
}

/*
 * @title: GetNodeInfo
 * @description: 获取Node相关信息
 * @param:
 * @return:
 */

func GetNodeInfo() *host.InfoStat {
	nodeInfo, _ := host.Info()
	return nodeInfo
}

/*
 * @title: GetDiskinfo
 * @description: 获取Disk相关信息
 * @param:
 * @return:
 */

func GetDiskinfo() []DiskInfo {
	diskInfo, _ := disk.Partitions(true)
	diskinfo := []DiskInfo{}
	for _, value := range diskInfo {
		diskusage, _ := disk.Usage(value.Device)
		diskinfo = append(diskinfo, DiskInfo{DiskInfo: value, Usage: *diskusage})
	}
	return diskinfo
}

func GetSysteminfo() NodeInfo {
	cpuInfo := GetCPUinfo()
	nodeInfo := GetNodeInfo()
	diskInfo := GetDiskinfo()
	systeminfo := NodeInfo{}
	systeminfo = NodeInfo{NODE: nodeInfo, CPU: cpuInfo, DISK: diskInfo}
	return systeminfo
}
