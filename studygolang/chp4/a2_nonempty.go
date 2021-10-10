/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: nonempty.py
@Time: 2021-10-13 17:07
@Last_update: 2021-10-13 17:07
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
	data := []string{"one", "", "three"}
	fmt.Println(nonempty(data))
	fmt.Println(data)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}