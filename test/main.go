package main

import (
	"fmt"
	"golang-tools/common/datetime"
)



func datetime_test()  {
	
}
func main() {
	//common.FormatPrint()
	//fmt.Println(common.TimeStamp())

	//数组示例
	// te := [5]string {"ss","qq","ww","rr","tt"}
	// 定义一个 Map
	// aMap := make(map[string]string, 8)
	// aMap["nihao"] = "qqq"
	// aMap["sss"] = "sss"
	// fmt.Printf("%v", aMap)

	// 字符串比较
	// name := "nihao"
	// ss := "NIHAO"
	// a := strings.Compare(name,ss)
	// fmt.Printf("%d",a)

	// start := common.NowTime()
	// time.Sleep(5*time.Second)
	// end := common.NowTime()

	// var diffTime = common.CalculateRunTime(start,end)

	// common.ConversionTime(3600)
	fmt.Printf("%s", datetime.ConversionTime(7896))

}
