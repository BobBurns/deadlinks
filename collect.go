package main

import (
	"golang.org/x/net/html"
	"io"
	"strconv"
	"strings"
)

/*
collectLinks taken from Jack Danger Canty at https://github.com/JackDanger/collectlinks
 Thanks!

The MIT License (MIT)

Copyright (c) 2014 Jack Danger Canty

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

//bob modification:
// take out redundant link check
// return array of link text
// handle links without text

func collectLinks(httpBody io.ReadCloser) ([]string, []string) {
	links := []string{}
	// col := []string{}
	txts := []string{}
	page := html.NewTokenizer(httpBody)
	previous := false
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links, txts
		}
		if tokenType == html.TextToken && previous == true {
		  txts = append(txts, strings.TrimSpace(string(page.Text())))
		  previous = false
		}

		token := page.Token()
		// handle image links etc
		if previous == true && tokenType == html.EndTagToken && token.DataAtom.String() == "a" {
		  txts = append(txts, "")
		  previous = false
		}

		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					// don't need this now I'm checking all links
					//col = append(col, tl)
					// skip resolv
					// resolv(&links, col)
					links = append(links, tl)
					previous = true
				}

			}
		}

	}
}

func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

func check(sl []string, s string) bool {
	var check bool
	for _, str := range sl {
		if str == s {
			check = true
			break
		}
	}
	return check
}

func resolv(sl *[]string, ml []string) {
	for _, str := range ml {
		if check(*sl, str) == false {
			*sl = append(*sl, str)
		}
	}
}
