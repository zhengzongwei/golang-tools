package datetime_test

import (
	"fmt"
	"golang-tools/common/datetime"
	"math/rand"
	"testing"
)

func TestConversionTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		timeSec := rand.Int63n(10000)
		f := datetime.ConversionTime(timeSec)
		fmt.Printf("%v\n",f)
	}

}
