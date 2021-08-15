/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a7_dup2.py
@Time: 2021-08-15 15:08
@Last_update: 2021-08-15 15:08
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
	/*
	整体流程：
	1. 得到文件的名称
	2. 分别调用计算
	 */
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}

		countlines(f, counts)
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countlines(file *os.File, counts map[string]int)  {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}