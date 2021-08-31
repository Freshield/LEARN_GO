/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: authority.py
@Time: 2021-08-31 19:46
@Last_update: 2021-08-31 19:46
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
	"chp5/counters"
	"fmt"
)

func main() {
	counter := counters.New(10)
	fmt.Println(counter)
}