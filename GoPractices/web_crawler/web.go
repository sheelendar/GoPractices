package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

// Write a simple web crawler that, given a URL A, prints all URLs reachable from A on the same host and protocol.
// Testcase: https://crawl.opendoor.com/1.html

// Write a function that takes a URL and extracts all the hyperlinks

// To execute Go code, please declare a func main() in a package "main"
const baseURLWithHttps = "https://crawl.opendoor.com"
const baseURL = "crawl.opendoor.com"

var cache = make(map[string]bool)
var httpClint = http.DefaultClient

func main() {
	url := "https://crawl.opendoor.com/1.html"
	crawlerAllURL(url)
}

func crawlerAllURL(url1 string) {

	if len(url1) <= 0 {
		return
	}

	res, err := httpClint.Get(url1)
	if err != nil {
		fmt.Println("get error from url")
		return
	}
	urls := getURLDoc(res)
	urlSize := len(urls)
	var fullURLs []string

	for i := 0; i < urlSize; i++ {
		if !strings.Contains(urls[i], baseURL) {
			if _, ok := strings.CutPrefix(urls[i], "/"); ok {
				fullURLs = append(fullURLs, fmt.Sprintf("%s%s", baseURLWithHttps, urls[i]))
			}
		} else {
			fullURLs = append(fullURLs, urls[i])
		}
	}

	for i := 0; i < len(fullURLs); i++ {
		if _, ok := cache[fullURLs[i]]; !ok {
			cache[fullURLs[i]] = true
			crawlerAllURL(fullURLs[i])
			fmt.Print(fullURLs[i])
			fmt.Println()
		}
	}

}

func getURLDoc(res *http.Response) []string {
	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return parseHtml(doc)
}

func parseHtml(n *html.Node) []string {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, element := range n.Attr {
			if element.Key == "href" {
				links = append(links, element.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		moreLinks := parseHtml(c)
		for i := 0; i < len(moreLinks); i++ {
			links = append(links, moreLinks[i])
		}

	}
	return links
}
