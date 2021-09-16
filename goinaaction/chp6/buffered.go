/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: buffered.py
@Time: 2021-08-23 16:02
@Last_update: 2021-08-23 16:02
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

const (
	numberGoroutines = 4
	taskLoad = 10
)

var wg8 sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	wg8.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}
	fmt.Println("begin close")
	close(tasks)
	fmt.Println("tasks close")

	wg8.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg8.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Println("Worker:", worker, "Shutting Down")
			return
		}

		fmt.Println("Worker", worker, "Started", task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Println("Worker", worker, "Completed", task)
	}
}