/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: handler_test.py
@Time: 2021-09-27 13:11
@Last_update: 2021-09-27 13:11
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
	"net/http"
	"net/http/httptest"
	"chp9/handler"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	handler.Routes()
}

func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendJson", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request.",
				ballotX, err)
		}
		t.Log("\tShould be able to create a request.",
			checkMark)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark)

		u := struct {
			Name string
			Email string
		}{}
		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotX)
		}
		t.Log("\tShould decode the response.", checkMark)

		if u.Name == "Bill" {
			t.Log("\tShould have a Name.", checkMark)
		} else {
			t.Error("\tShould have a Name.", ballotX, u.Name)
		}

		if u.Email == "bill@studio.com" {
			t.Log("\tShould have an Email.", checkMark)
		} else {
			t.Error("\tShould have an Email.", ballotX, u.Email)
		}
	}
}

