/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a5_struct.py
@Time: 2021-10-13 22:03
@Last_update: 2021-10-13 22:03
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

import test "fmt"

type point struct {
	X int
	Y int
}

type Circle struct {
	point
	Radius int
}

func main() {
	circle := Circle{point: point{1, 2}, Radius: 3}
	test.Println(circle.X)
}