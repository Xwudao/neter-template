package utils

import (
	"math/rand"
	"strconv"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	// 定义随机字符串的字符集
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 生成随机字符串
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charSet[r.Intn(len(charSet))]
	}

	return string(randomString)
}

// GenRandNumber 生成指定倍数的随机数
func GenRandNumber(length int) string {
	// 定义随机字符串的字符集
	charSet := "0123456789"

	// 生成随机字符串
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charSet[r.Intn(len(charSet))]
	}

	return string(randomString)
}

// GenOrderID 生成订单号
func GenOrderID(uid int64) string {
	var now = time.Now()
	return now.Format("20060102150405") + strconv.FormatInt(uid, 10) + GenRandNumber(4)
}
