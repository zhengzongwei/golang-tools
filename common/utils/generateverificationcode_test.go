package utils

import (
	"fmt"
	"testing"
)

func TestGenerateVerificationCode(t *testing.T) {
	numberCode := GenerateVerificationCode("number", 6)
	fmt.Println(numberCode)

	letterCode := GenerateVerificationCode("letter", 6)
	fmt.Println(letterCode)

	letterNumberCode := GenerateVerificationCode("letter_number", 6)
	fmt.Println(letterNumberCode)
}
