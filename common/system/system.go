/*
 * @Title system.go
 * @Description 获取系统相关信息 CPU 网卡 系统类型等等信息
 * @Author zhengzongwei 2021/11/16 11:16
 * @Update
 */

package system

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func SystemInfo(){
	nodeInfo,_:= host.Info()
	
	fmt.Printf("%v",nodeInfo)
 }

