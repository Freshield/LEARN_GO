/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: handlers_example_test.py
@Time: 2021-09-27 13:47
@Last_update: 2021-09-27 13:47
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package handler_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendJson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)

	var u struct{
		Name string
		Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("ERROR:", err)
	}

	fmt.Println(u)
	// Output:
	// {Bill bill@studio.com}
}