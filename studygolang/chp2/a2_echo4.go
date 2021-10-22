/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a2_echo4.py
@Time: 2021-10-08 17:38
@Last_update: 2021-10-08 17:38
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
	"strings"
	"strconv"
)

func main() {
	n, sep := false, " "
	np, nsep := &n, &sep
	args := []string{strconv.FormatBool(n), *nsep}
	fmt.Print(strings.Join(args, sep))
	if !*np {
		fmt.Println()
	}
}


