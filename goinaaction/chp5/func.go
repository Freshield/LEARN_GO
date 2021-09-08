/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: func.py
@Time: 2021-09-09 16:15
@Last_update: 2021-09-09 16:15
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

import "fmt"

type user1 struct {
	name string
	email string
}

func (u user1) notify() {
	fmt.Println(u.name, u.email)
}

func (u *user1) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user1{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user1{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@comcast.com")
	lisa.notify()
}
