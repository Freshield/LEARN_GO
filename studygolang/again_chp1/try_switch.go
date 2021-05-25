/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: try_switch.py
@Time: 2021-09-30 18:02
@Last_update: 2021-09-30 18:02
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
	"math/rand"
	"time"
)

type Point struct {
	X, Y int
}

var p Point

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Float64()
	fmt.Println(num)
	var coin string
	switch {
	case num > 0.5:
		coin = "heads"
	case num < 0.5:
		coin = "tails"
	default:
		coin = "edge"
	}

	switch coin {
	case "heads":
		fmt.Println("heads")
	case "tails":
		fmt.Println("tails")
	default:
		fmt.Println("landed on edge!")
	}
}