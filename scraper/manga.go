package scraper

import (
	"fmt"
	"image"
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Page represents a page within the manga
type Page struct {
	Idx    int
	Height int
	Width  int
	// Img is the art on the page
	Img image.Image
}

// Manga represents a manga
type Manga struct {
	// The name of the manga
	Name string
	// The actual pages
	Pages []*Page
}

// GeneratePDF turns the Manga into a PDF document
func (m *Manga) GeneratePDF() {
}

// ParseManga takes in a URL
func ParseManga(r io.Reader, options *Options) *Manga {
	if options == nil {
		options = DefaultOptions
	}

	t := html.NewTokenizer(r)

	var pages []*Page
	var name string

	parsingChapters := false
	for {
		nxt := t.Next()

		if nxt == html.ErrorToken {
			break
		} else {
			tagName, hasAttr := t.TagName()

			switch atom.Lookup(tagName) {
			case atom.Img:
				{

				}
			case atom.Div:
				{

				}
				// TODO: This is kind of a hacky way to get the mangas title. I am probably just going to not bother with this.
			case atom.Select:
				{
					if hasAttr {
						for {
							key, value, hasMore := t.TagAttr()

							if string(key) == "name" && string(value) == "chapterMenu" {
								fmt.Println(string(t.Raw()))
								parsingChapters = true
								fmt.Println("Parsing chapters")
							}

							if !hasMore {
								break
							}
						}
					}
				}
			case atom.Option:
				{
					if hasAttr {
						for {
							key, value, hasMore := t.TagAttr()
							if string(key) == "selected" && string(value) == "selected" && parsingChapters {
								nxt = t.Next()
								if nxt == html.TextToken {
									fmt.Println(t.Text())
								}
							}

							if !hasMore {
								break
							}
						}
					}
				}
			}

		}
	}

	return &Manga{
		Name:  name,
		Pages: pages,
	}
}

// NewPageURL creates a new page from a URL
func NewPageURL(url string, idx int) *Page {
	return nil
}
