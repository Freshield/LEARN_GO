/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: test.py
@Time: 2021-08-20 17:02
@Last_update: 2021-08-20 17:02
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
	"a1_search_match/search"
	"fmt"
)

func main() {
	feeds, err := search.RetrieveFeeds()
	if err != nil {
		fmt.Println(err)
	}
	for _, feed := range feeds {
		fmt.Println(feed.URI)
	}
}