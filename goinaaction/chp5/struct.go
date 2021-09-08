/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: struct.py
@Time: 2021-09-10 16:08
@Last_update: 2021-09-10 16:08
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

type user struct {
	name string
	email string
	ext int
	privileged bool
}

type admin struct {
	person user
	level string
}

func main() {
	var bill user

	lisa := user{
		name: "Lisa", email: "lisa@email.com",
		ext: 123, privileged: true,
	}
	lisa1 := user{"Lisa", "lisa@email.com", 123, true}
	fmt.Println(bill)
	fmt.Println(lisa)
	fmt.Println(lisa1)

	fred := admin{
		person: user{
			name: "Lisa", email: "lisa@email.com",
			ext: 123, privileged: true,
		},
		level: "super",
	}
	fmt.Println(fred)

}