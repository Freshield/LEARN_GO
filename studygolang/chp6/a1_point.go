/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a1_point.py
@Time: 2021-10-22 13:12
@Last_update: 2021-10-22 13:12
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

type Point struct {
	X float64
	Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p *Point) isNil() {
	if p == nil {
		fmt.Println("is nil")
	}
}

func main() {
	p := Point{1, 2}
	fmt.Println(p)
	p.ScaleBy(2)
	fmt.Println(p)
	p.isNil()

	p1 := &Point{}
	p1.isNil()
	fmt.Println(p1)
	p1 = nil
	p1.isNil()

}