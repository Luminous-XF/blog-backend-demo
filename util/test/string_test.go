package test

import (
	"blog-backend/util"
	"fmt"
	"testing"
)

func TestCamelToSnake(t *testing.T) {
	camelStr := ""
	fmt.Printf("'%s' -> '%s'\n", camelStr, util.CamelToSnake(camelStr))

	camelStr = "User"
	fmt.Printf("'%s' -> '%s'\n", camelStr, util.CamelToSnake(camelStr))

	camelStr = "UserId"
	fmt.Printf("'%s' -> '%s'\n", camelStr, util.CamelToSnake(camelStr))

	camelStr = "USERid"
	fmt.Printf("'%s' -> '%s'\n", camelStr, util.CamelToSnake(camelStr))

	camelStr = "USER"
	fmt.Printf("'%s' -> '%s'\n", camelStr, util.CamelToSnake(camelStr))

	camelStr = "user"
	fmt.Printf("'%s' -> '%s'\n", camelStr, util.CamelToSnake(camelStr))

	camelStr = "ThisIsCamelCase"
	fmt.Printf("'%s' -> '%s'\n", camelStr, util.CamelToSnake(camelStr))

	camelStr = "This_IsCamelCase"
	fmt.Printf("'%s' -> '%s'\n", camelStr, util.CamelToSnake(camelStr))

	camelStr = "This_Is_CamelCase"
	fmt.Printf("'%s' -> '%s'\n", camelStr, util.CamelToSnake(camelStr))
}
