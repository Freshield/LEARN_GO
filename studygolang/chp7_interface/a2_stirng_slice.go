/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a2_stirng_slice.py
@Time: 2021-10-26 13:27
@Last_update: 2021-10-26 13:27
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
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	names := []string{"a", "c", "d", "b", "z", "f"}
	fmt.Println(names)
	sort.Sort(StringSlice(names))
	fmt.Println(names)
}