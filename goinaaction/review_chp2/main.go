/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: main.py
@Time: 2021-09-29 11:19
@Last_update: 2021-09-29 11:19
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
	"log"
	"os"

	_ "review_chp2/matchers"
	"review_chp2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}