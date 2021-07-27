package common

import (
	"fmt"
	"time"
)

func FormatPrint() {
	fmt.Printf("this is a common.go")
}

func TimeStamp() int64 {
	currentTimeStamp := time.Now().Unix()
	return currentTimeStamp
}

func NowTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime
}

//TODO 创建一个记录程序运行时间的函数
