package main

import (
	"log"
	"net/http"

	"github.com/rbrick/mangapanda-scraper/scraper"
)

func main() {
	r, err := http.Get("http://www.mangapanda.com/fairy-tail/495")
	if err != nil {
		log.Fatalln(err)
	}
	scraper.ParseManga(r.Body, nil)
}
