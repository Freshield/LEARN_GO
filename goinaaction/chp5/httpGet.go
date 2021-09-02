/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: listing34.py
@Time: 2021-09-07 16:32
@Last_update: 2021-09-07 16:32
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
	"io"
	"net/http"
	"os"
)

func main() {
	url := "http://pdclound.medai.qq.com:8090"
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}