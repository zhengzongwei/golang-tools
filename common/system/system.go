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
	diskinfo disk.PartitionStat
	usage disk.UsageStat
}

type NodeInfo struct{
	os *host.InfoStat
	cpu[] cpu.InfoStat
	disk DiskInfo
}

/*
 * @title: SystemInfo
 * @description: 获取系统相关信息
 * @param:
 * @return:
 */

func SystemInfo(){
	nodeInfo,_:= host.Info()
	cpuInfo, _:= cpu.Info()

	diskInfo,_:=disk.Partitions(true)
	//for i:=0;i<len(diskInfo);i++{
	//	fmt.Printf("%v\n",diskInfo[i])
	//}
	var diskinfo DiskInfo = DiskInfo{ }
	for _,value :=range diskInfo{


		diskusage,_ := disk.Usage(value.Device)
		diskinfo.diskinfo = value
		diskinfo.usage = *diskusage
	}


	info :=NodeInfo{
		os:nodeInfo,
		cpu:cpuInfo,
		disk:diskinfo,
	}
	
	fmt.Printf("%v\n",diskinfo)
	//fmt.Printf("%v\n",diskinfo.usage)
	fmt.Printf("%v\n",info.cpu)
	fmt.Printf("%v\n",info.disk)
 }

