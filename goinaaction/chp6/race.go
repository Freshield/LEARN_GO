/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: race.py
@Time: 2021-08-28 14:29
@Last_update: 2021-08-28 14:29
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
)

var (
	counter int
	wg1 sync.WaitGroup
)

func main() {
	wg1.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg1.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg1.Done()

	for count := 0; count < 2; count++ {
		value := counter

		runtime.Gosched()

		value++

		counter = value
	}
}