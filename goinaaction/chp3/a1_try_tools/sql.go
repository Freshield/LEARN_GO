/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: sql.py
@Time: 2021-08-27 18:03
@Last_update: 2021-08-27 18:03
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
	"database/sql"
	_ "github.com/goinaction/code/chapter3/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
}