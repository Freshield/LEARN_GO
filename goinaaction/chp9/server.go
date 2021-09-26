/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: server.py
@Time: 2021-09-27 13:02
@Last_update: 2021-09-27 13:02
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
	"net/http"
	"chp9/handler"
)

func main() {
	handler.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}


