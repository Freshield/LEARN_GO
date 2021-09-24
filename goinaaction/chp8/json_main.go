/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: json_main.py
@Time: 2021-09-24 14:00
@Last_update: 2021-09-24 14:00
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
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Contact struct{
		Home string `json:"home"`
		Cell string `json:"cell"`
	}
}

var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(c)

	var d map[string]interface{}
	err = json.Unmarshal([]byte(JSON), &d)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(d)
	fmt.Println(d["contact"])
	fmt.Println(d["contact"].(map[string]interface{})["home"])
	e := d["contact"].(map[string]interface{})["home"]
	fmt.Printf("%T", e)

	f := make(map[string]interface{})
	f["name"] = "Gopher"
	f["title"] = "progammer"
	f["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	data, err := json.MarshalIndent(f, "", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))
	fmt.Println(data)
}



