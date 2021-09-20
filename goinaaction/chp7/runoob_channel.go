/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: runoob_channel.py
@Time: 2021-09-23 15:59
@Last_update: 2021-09-23 15:59
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

func fibonacci(c, quit chan int) {
	fmt.Println("in fib")
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("c <- x", "x", x, "y", y)
			x, y = y, x+y
			time.Sleep(time.Second)
			fmt.Println("done sleep")
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("in func()")
		for i := 0; i < 10; i++ {
			fmt.Println("i", i)
			fmt.Println(<-c)
		}
		fmt.Println("done to quit")
		quit <- 0
	}()
	fmt.Println("ready")
	fibonacci(c, quit)
}