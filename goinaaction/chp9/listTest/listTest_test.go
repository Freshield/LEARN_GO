/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: listTest_test.py
@Time: 2021-09-26 15:17
@Last_update: 2021-09-26 15:17
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package listTest

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	var urls = []struct{
		url string
		statusCode int
	}{
		{
			"http://www.goinggo.net/index.xml",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
				u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to make the Get call.",
						ballotX, err)
				}
				t.Log("\t\tShould be able to be able to make the Get call,",
					checkMark)

				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould receive a \"%d\" status. %v",
						u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
						u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}

	}
}

