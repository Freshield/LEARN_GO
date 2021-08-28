/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: sliceTime.py
@Time: 2021-09-16 14:11
@Last_update: 2021-09-16 14:11
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
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

func printPrime(prefix string) {
	defer wg.Done()

	next:
		for outer := 2; outer < 5000; outer++ {
			for inner :=2; inner < outer; inner++ {
				if outer % inner == 0 {
					continue next
				}
			}
			fmt.Println(prefix, outer)
		}
		fmt.Println("completed", prefix)

}