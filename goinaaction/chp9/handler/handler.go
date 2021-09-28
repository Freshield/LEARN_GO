/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: handler.py
@Time: 2021-09-27 13:03
@Last_update: 2021-09-27 13:03
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package handler

import (
	"encoding/json"
	"net/http"
)

func Routes() {
	http.HandleFunc("/sendJson", SendJSON)
}

func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name string
		Email string
	}{
		Name: "Bill",
		Email: "bill@studio.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}


