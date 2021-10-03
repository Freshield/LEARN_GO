/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a4_get_elements.py
@Time: 2021-10-19 19:54
@Last_update: 2021-10-19 19:54
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
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	url := "https://golang.org"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("resp error", err)
		return
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("doc error", err)
		return
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}