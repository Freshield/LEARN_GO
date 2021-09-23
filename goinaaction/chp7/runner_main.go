/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: runner_main.py
@Time: 2021-09-22 17:27
@Last_update: 2021-09-22 17:27
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
	"chp7/runner"
	"log"
	"os"
	"time"
)

const timeout = 2 * time.Second

func main() {
	log.Println("Starting work.")
	
	r := runner.New(timeout)
	
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)

		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}

	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Println("Processor - Task #", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

