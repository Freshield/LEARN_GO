/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: io_test.py
@Time: 2021-08-27 20:30
@Last_update: 2021-08-27 20:30
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
	"github.com/goinaction/code/chapter3/words"
	"io/ioutil"
)

func main() {
	filename := "sql.go"

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}