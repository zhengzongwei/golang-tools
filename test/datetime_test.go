package main

import (
	"fmt"
	"golang-tools/common/datetime"
	"testing"
)

func TestConversionTime(t *testing.T){
	f := datetime.ConversionTime(2046)
	fmt.Printf("%s",f)
	
}