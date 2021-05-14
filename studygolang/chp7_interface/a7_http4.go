/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a7_http4.py
@Time: 2021-10-26 22:59
@Last_update: 2021-10-26 22:59
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

type dollars3 float32

func (d dollars3) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database3 map[string]dollars3

func (db database3) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database3) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database3{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("0.0.0.0:9666", nil))
}