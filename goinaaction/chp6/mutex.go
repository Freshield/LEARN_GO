/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: mutex.py
@Time: 2021-08-26 15:07
@Last_update: 2021-08-26 15:07
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

var (
	counter5 int
	wg5 sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg5.Add(2)

	go incCounter5(1)
	go incCounter5(2)

	wg5.Wait()
	fmt.Println("Final Counter", counter5)
}

func incCounter5(id int) {
	defer wg5.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			fmt.Println(id, "here")
			value := counter5
			runtime.Gosched()
			value++
			counter5 = value
			fmt.Println(id, "gone")
		}
		mutex.Unlock()
	}

}