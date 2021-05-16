/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a5_http2.py
@Time: 2021-10-26 19:38
@Last_update: 2021-10-26 19:38
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
	"log"
	"net/http"
)

type dollars1 float32

func (d dollars1) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database1 map[string]dollars1

func (db database1) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func main() {
	db := database1{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("0.0.0.0:9666", db))
}