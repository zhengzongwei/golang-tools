package utils

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	path := "../conf/demo.yaml"
	conf := GetConfig(path)
	fmt.Printf("%#v", conf)
}
