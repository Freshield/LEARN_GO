/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: embedding_overwrite.py
@Time: 2021-08-31 19:34
@Last_update: 2021-08-31 19:34
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

type notifier7 interface {
	notify()
}

type user7 struct {
	name string
	email string
}

func (u *user7) notify() {
	fmt.Println(u.name, u.email)
}

type admin7 struct {
	user7
	level string
}

func (a *admin7) notify() {
	fmt.Println("admin", a.name, a.email)
}

func main() {
	ad := admin7{
		user7: user7{
			name: "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	sendNotification7(&ad)

	ad.user7.notify()

	ad.notify()
}

func sendNotification7(n notifier7) {
	n.notify()
}


