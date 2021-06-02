/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: dup2.py
@Time: 2021-09-30 11:22
@Last_update: 2021-09-30 11:22
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
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := []string{"test.txt", "test1.txt", "test2.txt"}

	for _, file_path := range files {
		f, err := os.Open(file_path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}
	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}