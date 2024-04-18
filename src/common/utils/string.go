package utils

import (
	"math/rand"
	"siteol.com/smart/src/common/constant"
	"strings"
	"time"
)

var baseStr = "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
var traceStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// TraceID 生成8位随机日志ID
func TraceID() string {
	return RandStr(9, true)
}

// SaltKey 生成一个16位的随机盐值
func SaltKey() string {
	return RandStr(16, false)
}

// RandStr 生成指定位数的随机字符
// f=true:大小英数字 =false:大写英文字母
func RandStr(length int, f bool) string {
	bytes := []byte(baseStr)
	if f {
		bytes = []byte(traceStr)
	}
	result := make([]byte, length)
	// Since 1.17+
	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(10000))))
	for i := 0; i < length; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

// StatusBool 状态转bool，0为true
func StatusBool(status string) bool {
	return status == constant.StatusOpen
}

// ArrayStr 字符串转数组
func ArrayStr(str string) []string {
	array := make([]string, 0)
	if str == "" {
		return array
	}
	// 逗号截取
	return strings.Split(str, ",")
}
