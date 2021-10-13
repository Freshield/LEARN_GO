/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a3_bool_int_transfer.py
@Time: 2021-10-09 20:45
@Last_update: 2021-10-09 20:45
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

func main() {
	fmt.Println(btoi(false))
	fmt.Println(itob(1))
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {return i != 0}