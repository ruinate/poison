package main

import (
	"PoisonFlow/src/config"
	"os"
)

// main 主执行程序
func main() {
	err := config.Poison.Execute()
	if err != nil {
		os.Exit(1)
	}
}
