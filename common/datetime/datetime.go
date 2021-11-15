/*
 * @Title common.go
 * @Description
 * @Author zhengzongwei 2021/11/10 10:16
 * @Update
 */

package datetime

import (
	"fmt"
	"strconv"
	"time"
)

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
 * @return: string
 */
func ConversionTime(timeStamp int64) string {

	days := strconv.FormatInt(timeStamp/(3600*24), 10)
	hour := strconv.FormatInt((timeStamp%(3600*24))/3600, 10)
	min := strconv.FormatInt((timeStamp%3600)/60, 10)
	sec := strconv.FormatInt(timeStamp&60/1000, 10)
	return days + "天" + hour + "时" + min + "分" + sec + "秒"
}
