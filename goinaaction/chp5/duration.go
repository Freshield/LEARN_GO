/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: duration.py
@Time: 2021-09-04 17:09
@Last_update: 2021-09-04 17:09
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

type duration int

func (d *duration) pretty() string {
	fmt.Println(*d)
	return fmt.Sprint(*d)
}

func main() {
	d := duration(42)
	fmt.Println(d)
	fmt.Println(&d)

	point := &d
	rst := point.pretty()
	fmt.Println(rst)
}

