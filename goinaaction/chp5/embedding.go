/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: embedding.py
@Time: 2021-09-02 19:22
@Last_update: 2021-09-02 19:22
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

type user5 struct {
	name string
	email string
}

func (u *user5) notify() {
	fmt.Println(u.name, u.email)
}

type admin5 struct {
	user5
	level string
}

func main() {
	ad := admin5{
		user5: user5{
			name: "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	ad.user5.notify()

	ad.notify()
}


