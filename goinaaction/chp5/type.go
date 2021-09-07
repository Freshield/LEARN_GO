/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: type.py
@Time: 2021-09-08 16:25
@Last_update: 2021-09-08 16:25
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

import "fmt"

func isShellSpecialVar(c uint8) bool {
	switch c {
	case '*', '#', '$':
		return true
	}
	return false
}

func main() {
	fmt.Println(isShellSpecialVar('1'))
}