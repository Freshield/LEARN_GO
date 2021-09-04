/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: funcSet.py
@Time: 2021-09-05 17:03
@Last_update: 2021-09-05 17:03
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

type notifier interface {
	notify()
}

type user2 struct {
	name string
	email string
}

func (u *user2) notify()  {
	fmt.Println(u.name, u.email)
}

func main() {
	u := user2{"Bill", "bill@email.com"}
	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}