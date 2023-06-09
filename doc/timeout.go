// Package doc -----------------------------
// @file      : timeout.go
// @author    : fzf
// @time      : 2023/4/20 上午11:53
// -------------------------------------------
package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.After(time.Second * 10)

	finish := make(chan bool)
	count := 1
	go func() {
		for {
			select {
			case <-timeout:
				fmt.Println("timeout")
				finish <- true
				return
			default:
				fmt.Printf("haha %d\n", count)
				count++
			}
			time.Sleep(time.Second * 1)
		}
	}()

	<-finish

	fmt.Println("Finish")
}
