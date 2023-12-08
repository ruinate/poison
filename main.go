// Package poison -----------------------------
// @file      : main.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午12:46
// -------------------------------------------
package main

import (
	logger "github.com/sirupsen/logrus"
	"poison/src/cmd"
)

// main 主执行程序
func main() {
	if err := cmd.PoisonCmd.Execute(); err != nil {
		logger.Fatalln(err)
	}
}
