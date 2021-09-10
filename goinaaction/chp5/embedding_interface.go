/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: embedding_interface.py
@Time: 2021-09-10 19:29
@Last_update: 2021-09-10 19:29
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
)

type notifier6 interface {
	notify()
}

type user6 struct {
	name string
	email string
}

func (u *user6) notify() {
	fmt.Println(u.name, u.email)
}

type admin6 struct {
	user6
	level string
}

func main() {
	ad := admin6{
		user6: user6{
			name: "john smith",
			email: "joh@yahoo.com",
		},
		level: "super",
	}

	sendNotification6(&ad)
}

func sendNotification6(n notifier6) {
	n.notify()
}

