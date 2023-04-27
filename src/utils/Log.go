// Package utils -----------------------------
// @file      : Log.go
// @author    : fzf
// @time      : 2023/4/20 上午10:57
// -------------------------------------------
package utils

import "log"

func LogError(err error) {
	log.Println(err.Error())
	return
}
