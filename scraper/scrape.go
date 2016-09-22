package scraper

// DefaultOptions are the default options used if none are provided.
var DefaultOptions = &Options{
	StartAt: 1,
	Name:    "Manga",
	BaseURL: "http://www.mangapanda.com",
}

// Options contain all the options for the scraper
type Options struct {
	// The index at which the scraper should start at.
	StartAt int
	// Manually set the name.
	Name    string
	BaseURL string
}
