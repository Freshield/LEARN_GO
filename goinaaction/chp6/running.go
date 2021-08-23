/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: running.py
@Time: 2021-08-24 15:45
@Last_update: 2021-08-24 15:45
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
	"time"
)

var wg7 sync.WaitGroup

func main() {
	baton := make(chan int)

	wg7.Add(1)

	go Runner(baton)

	baton <- 1

	wg7.Wait()
}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton

	fmt.Println("Runner", runner, "Running")

	if runner != 4 {
		newRunner = runner + 1
		fmt.Println("Runner", runner, "To The Line")
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Println("Runner", runner, "Finished")
		wg7.Done()
		return
	}

	fmt.Println("Runner", runner, "Exchange With Runner", newRunner)

	baton <- newRunner
}
