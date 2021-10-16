/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: popcount.py
@Time: 2021-10-09 15:01
@Last_update: 2021-10-09 15:01
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

var pc [256]byte

func init() {
	fmt.Println(pc)
	for i := range pc {
		fmt.Println("i", i)
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main()  {
	fmt.Println(pc)
}