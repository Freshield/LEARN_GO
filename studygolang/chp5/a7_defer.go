/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a7_defer.py
@Time: 2021-10-20 15:56
@Last_update: 2021-10-20 15:56
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

func trace(msg string) func() {
	start := time.Now()
	fmt.Println("enter", msg)
	return func() {
		fmt.Println("exit", msg, time.Since(start))
	}
}

func main() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}