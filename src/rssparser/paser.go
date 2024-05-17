package rssparser

import (
	"log"

	"github.com/mmcdole/gofeed"
)

const (
	blogRssFeed    = "https://cwc1222.github.io/rss.xml"
	maxPostsToShow = 5
)

func Parse() *gofeed.Feed {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(blogRssFeed)
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	return feed
}
