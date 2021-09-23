/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: runoob_keep_in_step.py
@Time: 2021-09-17 16:40
@Last_update: 2021-09-17 16:40
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
	"time"
)

func worker(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
	fmt.Println("done")
}



