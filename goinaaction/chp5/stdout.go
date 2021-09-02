/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: listing35.py
@Time: 2021-09-06 16:46
@Last_update: 2021-09-06 16:46
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
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("Hello"))

	fmt.Fprintf(&b, "World!")

	io.Copy(os.Stdout, &b)
}
