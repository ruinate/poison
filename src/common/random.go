// Package common -----------------------------
// @file      : random.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午12:50
// -------------------------------------------
package common

import (
	"math/rand"
	"time"
)

// RandStr 生成随机字符串
func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
