/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: runoob_close.py
@Time: 2021-09-18 16:34
@Last_update: 2021-09-18 16:34
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

func main() {
	go func() {
		time.Sleep(time.Hour)
	}()
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	//for i := range c {
	//	fmt.Println(i)
	//}
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	i, ok := <-c
	fmt.Println("i", i, "ok", ok)
}


