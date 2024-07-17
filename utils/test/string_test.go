package test

import (
	"blog-backend/utils"
	"fmt"
	"testing"
)

func TestGenerateString(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var str string
		str = utils.MakeStr(16, utils.Digit)
		fmt.Println(str)

		str = utils.MakeStr(16, utils.Alpha)
		fmt.Println(str)

		str = utils.MakeStr(6, utils.DigitAlpha)
		fmt.Println(str)

		str = utils.MakeStr(16, utils.DigitAlphaPunct)
		fmt.Println(str)
	}
}
