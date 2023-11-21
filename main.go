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
