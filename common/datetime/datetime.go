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
 * @Title datetime.go
 * @Description 时间处理模块
 * @Author zhengzongwei 2021/11/10 10:16
 * @Update
 */

package datetime

import (
	"fmt"
	"time"
)

func TimeStamp2TimeStr(timestamp int64) string {

	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

func TimeStr2TimeStamp(timeStr string) int64 {
	// Parse the time string
	Time, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		// Handle parsing error, for simplicity, returning 0
		fmt.Println("Error parsing time:", err)
		return 0
	}

	return Time.Unix()
}

/*
 * @title FormatPrint
 * @description
 * @param
 * @return None
 */

func FormatPrint() {
	fmt.Printf("this is a common.go")
}

/*
 * @title TimeStamp()
 * @description Convert time to timestamp
 * @param None
 * @return currentTimeStamp int64 "timestamp"
 */

func TimeStamp() int64 {
	currentTimeStamp := time.Now().Unix()
	return currentTimeStamp
}

/*
 * @title NowTime
 * @description get current time
 * @param None
 * @return currentTime string "time"
 */

func NowTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime
}

/*
 * @title: CalculateRunTime
 * @description: 计算程序的运行时间
 * @param:  startTime string
 *			endTime string
 * @return: int64 返回执行的时间
 */

func CalculateRunTime(startTime, endTime string) int64 {
	var diffTime int64
	timeTmp := "2006-01-02 15:04:05"
	starttime, err := time.ParseInLocation(timeTmp, startTime, time.Local)
	endtime, err := time.ParseInLocation(timeTmp, endTime, time.Local)
	if err != nil {
		return diffTime
	}
	diffTime = endtime.Unix() - starttime.Unix()
	return diffTime

}

/*
 * @title: CalculateRunTime
 * @description: 转换时间为具体的天周日时分秒
 * @param:  timeStamp int64 需要转换的时间戳
 * @return:  map[string]int64
 */

func ConversionTime(timeStamp int64) map[string]int64 {
	days := timeStamp / (3600 * 24)
	hour := (timeStamp % (3600 * 24)) / 3600
	_min := (timeStamp % 3600) / 60
	sec := timeStamp & 60 / 1000
	data := map[string]int64{"days": days, "hour": hour, "min": _min, "sec": sec}
	return data
}
