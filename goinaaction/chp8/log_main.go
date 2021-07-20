/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: log_main.py
@Time: 2021-09-23 17:01
@Last_update: 2021-09-23 17:01
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
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("message")

	//log.Fatalln("fatal message")

	log.Panicln("panic message")
}


