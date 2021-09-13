/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: slice.py
@Time: 2021-09-14 14:39
@Last_update: 2021-09-14 14:39
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
	slice := make([]string, 5)
	fmt.Println(slice)

	slice1 := make([]int, 3, 5)
	fmt.Println(slice1)

	slice2 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice2)
	slice3 := []int{10, 20, 30}
	fmt.Println(slice3)

	slice4 := []string{9: "1"}
	fmt.Println(slice4)

	var slice5 []int
	//slice5 = make([]int, 3)
	fmt.Println(slice5)
	slice6 := make([]int, 0)
	fmt.Println(slice6)
	slice7 := []int{}
	fmt.Println(slice7)

	slice8 := []int{10, 20, 30, 40, 50}
	slice8[1] = 25
	fmt.Println(slice8)

	newSlice := slice8[1: 3]
	fmt.Println(newSlice)
	newSlice[1] = 35
	fmt.Println(slice8)
	fmt.Println(newSlice)
	newSlice = append(newSlice, 60)
	fmt.Println(slice8)
	fmt.Println(newSlice)

	slice9 := []int{10, 20, 30, 40}
	newSlice1 := append(slice9, 50)
	fmt.Println(slice9)
	fmt.Println(newSlice1)

	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice10 := source[2:3:4]
	fmt.Println(source)
	fmt.Println(slice10)

	slice10 = append(slice10, "Kiwi")
	fmt.Println(source)
	fmt.Println(slice10)

	slice11 := source[2:3:3]
	slice11 = append(slice11, "Mango")
	fmt.Println(source)
	fmt.Println(slice11)

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Println(append(s1, s2...))

	for index, value := range slice11 {
		fmt.Println(index, value)
	}
	for index, value := range slice11 {
		fmt.Println(value, &value, slice11[index])
	}
	for _, value := range slice11 {
		fmt.Println(value)
	}
	for index := 1; index < len(slice11); index++ {
		fmt.Println(index, slice11[index])
	}

	slice12 := [][]int{{10}, {100, 200}}
	fmt.Println(slice12)
	slice12[0] = append(slice12[0], 20)
	fmt.Println(slice12)

}