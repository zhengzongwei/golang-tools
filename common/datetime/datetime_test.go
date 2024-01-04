package datetime

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestConversionTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		timeSec := rand.Int63n(10000)
		f := ConversionTime(timeSec)
		fmt.Printf("%v\n", f)
	}

}

func TestTimeStr2TimeStamp(t *testing.T) {
	TimeStamp := TimeStr2TimeStamp("2024-01-03 16:57:17")
	fmt.Printf("%d", TimeStamp)
}
func TestTimeStamp2TimeStr(t *testing.T) {
	TimeStr := TimeStamp2TimeStr(1704272237)
	fmt.Printf("%s", TimeStr)
}
