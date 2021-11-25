/*
 * @Title system.go
 * @Description 获取系统相关信息 CPU 网卡 系统类型等等信息
 * @Author zhengzongwei 2021/11/16 11:16
 * @Update
 */

package system

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

type DiskInfo struct {
	DiskInfo disk.PartitionStat `json:"diskInfo"`
	Usage    disk.UsageStat     `json:"usage"`
}

type NodeInfo struct {
	OS   *host.InfoStat `json:"os"`
	CPU  []cpu.InfoStat `json:"cpu"`
	DISK []DiskInfo     `json:"disk"`
}

/*
 * @title: SystemInfo
 * @description: 获取系统相关信息
 * @param:
 * @return:
 */

func SystemInfo() {
	nodeInfo, _ := host.Info()
	cpuInfo, _ := cpu.Info()

	diskInfo, _ := disk.Partitions(true)
	//for i:=0;i<len(diskInfo);i++{
	//	fmt.Printf("%v\n",diskInfo[i])
	//}
	diskinfo := []DiskInfo{}
	for _, value := range diskInfo {
		diskusage, _ := disk.Usage(value.Device)
		diskinfo = append(diskinfo, DiskInfo{diskinfo: value, usage: *diskusage})
	}

	info := NodeInfo{
		os:   nodeInfo,
		cpu:  cpuInfo,
		disk: diskinfo,
	}
	fmt.Printf("%v+\n", info)
	//fmt.Printf("%v\n",diskinfo.usage)
	//fmt.Printf("%v\n",info.cpu)
	//fmt.Printf("%v\n",info.disk)
}
