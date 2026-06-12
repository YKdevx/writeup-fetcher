package parser

import "encoding/xml"

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

func ParseRSS(data []byte) (*RSS, error) {
	var rss RSS

	err := xml.Unmarshal(data, &rss)
	if err != nil {
		return nil, err
	}

	return &rss, nil
}
