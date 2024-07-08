package util

import (
	"regexp"
	"strings"
)

func CamelToSnake(camelStr string) string {
	matched1, _ := regexp.MatchString(`([a-z])`, camelStr)
	matched2, _ := regexp.MatchString(`([A-Z])`, camelStr)
	if !matched1 || !matched2 {
		return strings.ToLower(camelStr)
	}

	re := regexp.MustCompile(`([a-z])([A-Z])`)
	snakeStr := re.ReplaceAllString(camelStr, "${1}_${2}")

	return strings.ToLower(snakeStr)
}
