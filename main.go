package main

import (
	"fmt"
	"log"

	"github.com/mmcdole/gofeed"
	//"github.com/cwc1222/cwc1222/src/rssparser"
)

const (
	blogRssFeed    = "https://cwc1222.github.io/rss.xml"
	maxPostsToShow = 5
)

func parseFeeds() *gofeed.Feed {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(blogRssFeed)
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	return feed
}

func main() {
	feed := parseFeeds()
	fmt.Println(feed)
}
