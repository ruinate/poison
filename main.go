package main

import (
	"PoisonFlow/src/strategy"
	"os"
)

// main 主执行程序
func main() {

	err := strategy.Poison.Execute()
	if err != nil {
		os.Exit(0)
	}
}
