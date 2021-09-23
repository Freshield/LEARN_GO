/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: worker_main.py
@Time: 2021-09-23 14:44
@Last_update: 2021-09-23 14:44
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
	"chp7/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
	index int
}

func (m *namePrinter) Task() {
	log.Println("index", m.index, "name", m.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	index := 0
	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
				index: index,
			}
			index++

			go func() {
				log.Println("i", i, "name", name)
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}


