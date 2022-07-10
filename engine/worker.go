package engine

import (
	"github.com/iralance/go-clawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("start url %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+
			"fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}
