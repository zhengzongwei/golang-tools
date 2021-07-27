package common

import (
	"fmt"
	"time"
)

func FormatPrint() {
	fmt.Printf("this is a common.go")
}

func NowTime() time.Time {
	currentTime := time.Now()
	return currentTime
}
