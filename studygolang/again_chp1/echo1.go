/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: echo1.py
@Time: 2021-09-29 14:55
@Last_update: 2021-09-29 14:55
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
	"os"
)

func main() {
	var s, sep string
	for i :=1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}