package test

import (
	"blog-backend/utils"
	"fmt"
	"testing"
)

func TestGenerateString(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var str string
		str = utils.MakeStr(16, utils.DIGIT)
		fmt.Println(str)

		str = utils.MakeStr(16, utils.ALAlpha)
		fmt.Println(str)

		str = utils.MakeStr(6, utils.ALNUM)
		fmt.Println(str)

		str = utils.MakeStr(16, utils.DigitAlphaPunct)
		fmt.Println(str)
	}
}
