package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: ", os.Args[0], "host port num")
		os.Exit(1)
	}

	target := os.Args[1]
	_ = os.Args[2]
	num, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid number of packets:", os.Args[3])
		os.Exit(1)
	}

	var wg sync.WaitGroup
	var lock sync.Mutex
	var pps int

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", target)
			if err != nil {
				return
			}
			conn.Close()

			lock.Lock()
			pps++
			lock.Unlock()
		}()
	}
	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Sent %d packets in %s (%.2f pps)\n", num, elapsed, float64(num)/elapsed.Seconds())
}
