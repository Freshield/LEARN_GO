/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: feed.py
@Time: 2021-08-20 23:14
@Last_update: 2021-08-20 23:14
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
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {return nil, err}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}