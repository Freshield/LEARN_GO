/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: array.py
@Time: 2021-09-15 14:06
@Last_update: 2021-09-15 14:06
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
	var array0 [5]int
	fmt.Println(array0)
	array := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array)
	array1 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(array1)
	array2 := [5]int{1: 10, 2: 20}
	fmt.Println(array2)

	array[2] = 35
	fmt.Println(array)

	array3 := [5]*int{0: new(int), 1: new(int)}
	*array3[0] = 10
	*array3[1] = 20
	fmt.Println(array3)

	var array4 [5]string
	array5 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	array4 = array5
	fmt.Println(array4)
	fmt.Println(array5)

	var array7 [3]*string
	array8 := [3]*string{new(string), new(string), new(string)}

	*array8[0] = "Red"
	*array8[1] = "Blue"
	*array8[2] = "Green"

	array7 = array8
	for _, point := range array7 {
		fmt.Println(*point)
	}

	//var array9 [4][2]int
	array9 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	array10 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	array11 := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array9)
	fmt.Println(array10)
	fmt.Println(array11)

	var array12 = array11[1]
	fmt.Println(array12)
	value := array11[1][1]
	fmt.Println(value)
	
}