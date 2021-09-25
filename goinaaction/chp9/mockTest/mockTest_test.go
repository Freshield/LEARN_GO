/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: mockTest_test.py
@Time: 2021-09-27 11:22
@Last_update: 2021-09-27 11:22
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package mockTest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
	<title>Going</title>
	<description>Golang</description>
	<link>goinggo.net</link>
	<item>
		<pubDate>sun</pubDate>
		<title>Object</title>
		<description>Go</description>
		<link>goinggo.net</link>
	</item>
</channel>
</rss>`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload(t *testing.T) {
	stausCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			server.URL, stausCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode != stausCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v",
					stausCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\tShould receive a \"%d\" status. %v",
				stausCode, checkMark)
		}
	}
}


