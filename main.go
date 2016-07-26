package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rbrick/mangapanda-scraper/scraper"
)

func main() {
	r, err := http.Get("http://www.mangapanda.com/soul-eater/1")

	if err != nil {
		log.Fatalln(err)
	}

	m := scraper.ParseManga(r.Body, nil)

	if m != nil {
		fmt.Println("Successfully parsed Manga")
		fmt.Println("Name:", m.Name)
	}
}
