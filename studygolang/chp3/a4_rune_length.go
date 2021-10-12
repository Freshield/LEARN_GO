/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a4_rune_length.py
@Time: 2021-10-09 20:53
@Last_update: 2021-10-09 20:53
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
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for _, text := range s {
		fmt.Printf("%q", text)
	}
}