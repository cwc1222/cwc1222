package main

import (
	"html/template"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

const (
	blogRssFeed    = "https://cwc1222.github.io/rss.xml"
	maxPostsToShow = 5

	readmeTmplPath = "README.md.tmpl"
	readmePath     = "README.md"
)

func parseFeeds(rssFeed string) *gofeed.Feed {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssFeed)
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
		panic(err)
	}
	return feed
}

func main() {
	feed := parseFeeds(blogRssFeed)

	tmpl, err := template.ParseFiles(readmeTmplPath)
	if err != nil {
		log.Fatalf("create file: %v", err)
		panic(err)
	}

	readme, err := os.Create(readmePath)
	if err != nil {
		log.Fatalf("create file: %v", err)
		panic(err)
	}
	defer readme.Close()

	n := min(maxPostsToShow, len(feed.Items))
	err = tmpl.Execute(readme, feed.Items[0:n])
	if err != nil {
		log.Fatalf("create file: %v", err)
		panic(err)
	}
}
