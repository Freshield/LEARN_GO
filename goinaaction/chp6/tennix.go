/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: tennix.py
@Time: 2021-08-25 15:24
@Last_update: 2021-08-25 15:24
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
	"sync"
	"time"
)

var wg6 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg6.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1

	wg6.Wait()
}

func player(name string, court chan int) {
	defer wg6.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Println("Player", name, "Won")
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Println("Player", name, "Missed")

			close(court)
			return
		}

		fmt.Println("Player", name, "Hit", ball)
		ball++

		court <- ball
	}
}


