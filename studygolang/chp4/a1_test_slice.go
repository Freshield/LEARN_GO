/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a1_test_slice.py
@Time: 2021-10-13 16:36
@Last_update: 2021-10-13 16:36
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

func main() {
	months := [...]int{1,2,3,4,5}
	fmt.Printf("%T\n", months)
	q1 := months[:4:4]
	fmt.Printf("%T\n", q1)
	fmt.Println(cap(q1))
	fmt.Println(q1)
	q1 = append(q1, 6)
	fmt.Println(q1)
	fmt.Println(months)

	var s1 []int
	s2 := []int(nil)
	s3 := []int{}
	s4 := make([]int, 3)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(s3 == nil)
	fmt.Println(s4)

}