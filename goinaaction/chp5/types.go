/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: types.py
@Time: 2021-09-02 17:40
@Last_update: 2021-09-02 17:40
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

type notifier1 interface {
	notify()
}

type user3 struct {
	name string
	email string
}

func (u *user3) notify() {
	fmt.Println(u.name, u.email)
}

type admin1 struct {
	name string
	email string
}

func (a *admin1) notify() {
	fmt.Println(a.name, a.email)
}

func main() {
	bill := user3{"Bill", "bill@email.com"}
	sendNotification1(&bill)

	lisa := admin1{
		name: "Lisa",
		email: "lisa@email.com",
	}
	sendNotification1(&lisa)
}

func sendNotification1(n notifier1) {
	n.notify()
}
