package test

import (
	"blog-backend/utils"
	"fmt"
	"testing"
)

func TestGenerateString(t *testing.T) {
	for i := 0; i < 10000; i++ {
		str := utils.GenerateSalt(16)
		fmt.Println(str)
	}
}
