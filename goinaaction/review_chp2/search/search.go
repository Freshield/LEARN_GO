/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: search.py
@Time: 2021-09-28 16:06
@Last_update: 2021-09-28 16:06
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()

		close(results)
	}()

	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}