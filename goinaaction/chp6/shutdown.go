/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: shutdown.py
@Time: 2021-08-27 14:54
@Last_update: 2021-08-27 14:54
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg3 sync.WaitGroup
)

func main() {
	wg3.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg3.Wait()
}

func doWork(name string) {
	defer wg3.Done()

	for {
		fmt.Println("Doing", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Println("Shutting", name)
			break
		}
	}
}

