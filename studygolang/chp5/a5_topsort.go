/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a5_topsort.py
@Time: 2021-10-19 20:18
@Last_update: 2021-10-19 20:18
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
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func visitAll(items []string, seen map[string]bool, m map[string] []string, order *[]string) {
	for _, item := range items {
		if !seen[item] {
			seen[item] = true
			visitAll(m[item], seen, m, order)
			*order = append(*order, item)
		}
	}
}

func topSort(m map[string][]string) []string {
	order := make([]string, 0, 100)
	seen := make(map[string]bool)
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys, seen, m, &order)
	return order
}

func main() {
	for i, course := range topSort(prereqs) {
		fmt.Println(i+1, "\t", course)
	}
}