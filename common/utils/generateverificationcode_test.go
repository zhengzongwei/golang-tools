package utils

import (
	"fmt"
	"testing"
)

func TestGenerateVerificationCode(t *testing.T) {
	code := GenerateVerificationCode("number", 6)
	fmt.Println(code)
}
