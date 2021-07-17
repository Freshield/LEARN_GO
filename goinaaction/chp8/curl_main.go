/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: curl_main.py
@Time: 2021-09-24 14:54
@Last_update: 2021-09-24 14:54
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
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("curl.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	dest := io.MultiWriter(os.Stdout, file)

	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

}

