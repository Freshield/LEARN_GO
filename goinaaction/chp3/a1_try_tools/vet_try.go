/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: vet_try.py
@Time: 2021-08-27 20:36
@Last_update: 2021-08-27 20:36
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

func main() {
	// fmt.Printf("The quick brown fox jumped over lazy dogs", 3.14)
	fmt.Printf("The quick brown fox jumped over lazy dogs %f", 3.14)
	err := "nil"
	if err != "nil" {
		return
	}
}
