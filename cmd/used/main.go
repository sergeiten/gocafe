package main

import (
	"flag"

	"github.com/sergeiten/gocafe/joonggonara/used"
)

var (
	query string
	pages int
	file  string
)

func init() {
	flag.StringVar(&query, "query", "", "Search query")
	flag.IntVar(&pages, "pages", 1, "Number pages for parsing")
	flag.StringVar(&file, "file", "", "Name of output xlsx file")

	flag.Parse()
}

func main() {
	list := used.Fetch(query, pages)
	used.WriteXlsFile(file, list)
}
