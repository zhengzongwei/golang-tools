package utils

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	path := "../conf/demo.yaml"
	conf := GetYamlConfig(path)
	fmt.Printf("%#v", conf)
}
