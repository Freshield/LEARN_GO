/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: word_link.py
@Time: 2021-09-24 14:41
@Last_update: 2021-09-24 14:41
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
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	fmt.Fprintf(&b, "World!")

	b.WriteTo(os.Stdout)

}


