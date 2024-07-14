package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	digitLetterBytes         = "0123456789"
	alphaLetterBytes         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitAndAlphaLetterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	allLetterBytes           = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ()`!@#$%^&*_-+=|{}[]:;'<>,.?"
)

var (
	letterIdxBits int64
	letterIdxMask int64
	letterIdxMax  int64
)

type MakeStrMode int

const (
	// ALPHA 仅包含大小写字母
	ALPHA MakeStrMode = iota
	// DIGIT 仅包含数字
	DIGIT
	// DIGIT_ALPHA 包含数字和大小写字母
	DIGIT_ALPHA
	// DIGIT_ALPHA_PUNCT 包含数字、字母、特殊符号
	DIGIT_ALPHA_PUNCT
)

var src = rand.NewSource(time.Now().UnixNano())

// MakeStr 生成包含大小写字母指定长度随机字符串
func makeStr(length int, letterBytes string) string {
	calMakeStrInfo()
	str := make([]byte, length)
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			str[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&str))
}

func calMakeStrInfo() {
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax = 63 / letterIdxBits
}

func MakeStr(length int, mode MakeStrMode) string {
	switch mode {
	case DIGIT:
		letterIdxBits = 4
		return makeStr(length, digitLetterBytes)
	case ALPHA:
		letterIdxBits = 6
		return makeStr(length, alphaLetterBytes)
	case DIGIT_ALPHA:
		letterIdxBits = 6
		return makeStr(length, digitAndAlphaLetterBytes)
	case DIGIT_ALPHA_PUNCT:
		letterIdxBits = 7
		return makeStr(length, allLetterBytes)
	default:
		letterIdxBits = 7
		return makeStr(length, allLetterBytes)
	}
}
