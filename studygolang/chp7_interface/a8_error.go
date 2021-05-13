/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a8_error.py
@Time: 2021-10-27 19:48
@Last_update: 2021-10-27 19:48
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
	"syscall"
)

func main() {
	var err = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}