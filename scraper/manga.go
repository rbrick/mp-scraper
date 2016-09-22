package scraper

import (
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Page represents a page within the manga
type Page struct {
	Idx int
	// Img is the art on the page
	Img image.Image
}

// Manga represents a manga
type Manga struct {
	// The name of the manga
	Name string
	// The actual pages
	Pages []string
}

// GeneratePDF turns the Manga into a PDF document
func (m *Manga) GeneratePDF() {

}

type parser struct {
	options  *Options
	pages    []string
	rootNode *html.Node
}

// ParseManga takes in a URL
func ParseManga(r io.Reader, options *Options) *Manga {
	if options == nil {
		options = DefaultOptions
	}

	p := &parser{
		options: options,
		pages:   []string{},
	}

	p.parse(r)
	return &Manga{
		Pages: p.pages,
	}
}

func (p *parser) parse(r io.Reader) {
	data, err := ioutil.ReadAll(r)

	if err != nil {
		log.Panicln(err)
	}

	htmls := string(data)
	rgx := regexp.MustCompile(`(document)(\[\'mangaid\'\])(\s+)(=)(\s+)(\d+)`)
	statement := rgx.FindString(htmls)
	id := rgx.ReplaceAllString(statement, "$6")
	fmt.Println("ID:", id)
	p.rootNode, err = html.Parse(strings.NewReader(htmls))

	p.gatherPages()

	fmt.Println(len(p.pages))
}

func (p *parser) gatherPages() {
	parseNode(p.rootNode, func(element *html.Node) {
		if element.DataAtom == atom.Option {
			attr := parseAttributes(element.Attr)
			if v, ok := attr["value"]; ok {
				parseURL(p.options.BaseURL+v, func(n *html.Node) {
					if n.DataAtom == atom.Img {
						attr := parseAttributes(n.Attr)
						if attr["id"] == "img" {
							p.pages = append(p.pages, saveImage(attr["src"]))
						}
					}
				})
			}
		}
	})
}

func parseURL(url string, f func(n *html.Node)) {
	r, err := http.Get(url)

	if err != nil {
		log.Panicln(err)
	}

	rn, err := html.Parse(r.Body)

	if err != nil {
		log.Panicln(err)
	}

	parseNode(rn, f)
}

func parseNode(n *html.Node, f func(n *html.Node)) {
parseLoop:
	for nxt := n.FirstChild; nxt != nil; nxt = nxt.NextSibling {
		switch nxt.Type {
		case html.ErrorNode:
			{
				break parseLoop
			}
		case html.ElementNode:
			{
				f(nxt)
			}
		}

		parseNode(nxt, f)
	}

}

func parseAttributes(a []html.Attribute) map[string]string {
	m := map[string]string{}
	for _, x := range a {
		m[x.Key] = x.Val
	}
	return m
}

// All these error checks man
func saveImage(url string) string {
	_, err := os.Stat("work")

	if os.IsNotExist(err) {
		err = os.Mkdir("work", os.ModeTemporary)

		if err != nil {
			log.Panicln(err)
		}
	}

	resp, err := http.Get(url)

	if err != nil {
		log.Panicln(err)
	}

	d, _ := ioutil.ReadAll(resp.Body)

	fname := filepath.Join("work", url[strings.LastIndex(url, "/"):])
	err = ioutil.WriteFile(fname, d, os.ModeTemporary)
	if err != nil {
		log.Panicln(err)
	}

	return fname
}
