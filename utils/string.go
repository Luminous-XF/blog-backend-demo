package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ()`!@#$%^&*_-+=|{}[]:;'<>,.?"

const (
	letterIdxBits = 7
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

// GenerateSalt 生成包含大小写字母指定长度随机字符串
func GenerateSalt(length int) string {
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
