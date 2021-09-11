/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: map.py
@Time: 2021-09-13 15:41
@Last_update: 2021-09-13 15:41
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
	dict := make(map[string]int)
	dict1 := map[string]string{"Red": "red", "Orange": "orange"}
	fmt.Println(dict)
	fmt.Println(dict1)

	dict2 := map[int] []string{}
	fmt.Println(dict2)

	colors := map[string]string{}
	colors["Red"] = "red"
	fmt.Println(colors)

	value, exists := colors["Blue"]
	fmt.Println(value, exists)
	value = colors["Blue"]
	fmt.Println(value)
	
	for key, value := range dict1 {
		fmt.Println(key, value)
	}
	delete(dict1, "Red")
	fmt.Println(dict1)
}