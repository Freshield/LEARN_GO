/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: default.py
@Time: 2021-08-22 01:52
@Last_update: 2021-08-22 01:52
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

type defaultMatcher struct {}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}