package main

import (
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 0)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		logger.Fatal("no JSON comics file")
	}

	if len(args) < 2 {
		logger.Fatal("no search arguments")
	}

	JSONStore, err := os.Open(args[0])
	if err != nil {
		logger.Fatal(err)
	}

	searcher, err := NewSearcher(JSONStore)
	if err != nil {
		logger.Fatal(err)
	}

	query := args[1:]
	matched := searcher.Search(query...)
	
	logger.Printf("read %v comics", searcher.Total())
	for _, comic := range matched {
		link := fmt.Sprintf("https://xkcd.com/%v/", comic.Number)
		date := fmt.Sprintf("%v/%v/%v", comic.Month, comic.Day, comic.Year)
		fmt.Println(link, date, comic.Title)
	}
	logger.Printf("found %v comics", len(matched))
}
