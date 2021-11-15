package main

import (
	"fmt"
	"golang-tools/common/datetime"
	"math/rand"
	"testing"
)

func TestConversionTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int63n(10000))
		f := datetime.ConversionTime(rand.Int63n(10000))
		fmt.Printf("%s\n", f)
	}

	// f := datetime.ConversionTime(2046)
	// fmt.Printf("%s",f)

}
