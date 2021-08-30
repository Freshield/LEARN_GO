/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: entities.py
@Time: 2021-08-30 19:49
@Last_update: 2021-08-30 19:49
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package entities

type user8 struct {
	Name string
	Email string
}

type Admin struct {
	user8
	Rights int
}