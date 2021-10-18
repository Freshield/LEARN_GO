/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a5_tempconv0.py
@Time: 2021-10-09 14:49
@Last_update: 2021-10-09 14:49
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
	"chp2/tempconv"
	"fmt"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.CToF(tempconv.FreezingC))

	c := tempconv.FToC(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(c.String())
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
}