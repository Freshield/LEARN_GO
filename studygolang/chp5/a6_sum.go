/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a6_sum.py
@Time: 2021-10-19 20:49
@Last_update: 2021-10-19 20:49
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
)

func sum(vals ...int) int {
	value := 0
	for _, val := range vals {
		value += val
	}
	return value
}

func main() {
	fmt.Println(sum(1,2,3,4,5))
	values := []int{1,2,3}
	fmt.Println(sum(values...))
}