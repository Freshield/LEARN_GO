/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a3_remove.py
@Time: 2021-10-13 17:10
@Last_update: 2021-10-13 17:10
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

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5,6,7,8}
	fmt.Println(remove(s, 2))
}