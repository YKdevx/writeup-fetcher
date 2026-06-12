package main

import (
	"fmt"

	"github.com/YKdevx/writeup-fetcher/internal/fetcher"
	"github.com/YKdevx/writeup-fetcher/internal/parser"
)

func main() {

	rssFetcher := fetcher.RSSFetcher{
		URL: "https://medium.com/feed/tag/cybersecurity",
	}

	data, err := rssFetcher.Fetch()
	if err != nil {
		panic(err)
	}

	rss, err := parser.ParseRSS(data)
	if err != nil {
		panic(err)
	}

	for _, item := range rss.Channel.Items {
		fmt.Println(item.Title)
	}
}
