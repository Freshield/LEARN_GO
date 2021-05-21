/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a2_isnil.py
@Time: 2021-10-22 13:17
@Last_update: 2021-10-22 13:17
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

type IntList struct {
	Value int
	Tail *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	list := IntList{1, &IntList{2, nil}}
	fmt.Println(list.Sum())
}