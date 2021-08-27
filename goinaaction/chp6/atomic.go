/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: atomic.py
@Time: 2021-08-28 14:46
@Last_update: 2021-08-28 14:46
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
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter2 int64
	wg2 sync.WaitGroup
)

func main() {
	wg2.Add(2)

	go incCounter2(1)
	go incCounter2(2)

	wg2.Wait()

	fmt.Println("Final Counter:", counter2)
}

func incCounter2(id int) {
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter2, 1)

		runtime.Gosched()
	}
}