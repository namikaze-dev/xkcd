package main

import (
	"encoding/json"
	"io"
	"strings"
)

type Searcher struct {
	comics []Comic
}

type Comic struct {
	Day    string `json:"day"`
	Month  string `json:"month"`
	Year   string `json:"year"`
	Number int    `json:"num"`
	Title  string `json:"title"`
	Transcript string `json:"transcript"`
}

func NewSearcher(store io.Reader) (*Searcher, error) {
	searcher := &Searcher{}

	err := json.NewDecoder(store).Decode(&searcher.comics)
	if err != nil {
		return nil, err
	}

	return searcher, nil
}

func (sch Searcher) Total() int {
	return len(sch.comics)
}

func (sch Searcher) Search(query ...string) []Comic {
	for i := 0; i < len(query); i++ {
		query[i] = strings.ToLower(query[i])
	}

	var matched []Comic

	for _, comic := range sch.comics {
		for _, q := range query {
			if strings.Contains(comic.Title, q) || strings.Contains(comic.Transcript, q) {
				matched = append(matched, comic)
				break
			}
		}
	}

	return matched
}
