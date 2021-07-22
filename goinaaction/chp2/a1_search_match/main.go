/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: main.py
@Time: 2021-08-19 19:20
@Last_update: 2021-08-19 19:20
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
	_ "a1_search_match/matchers"
	"a1_search_match/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}