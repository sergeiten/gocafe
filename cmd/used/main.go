package main

import (
	"flag"

	"github.com/sergeiten/gocafe/joonggonara/used"
)

func main() {
	query := flag.String("query", "", "Search Query")
	pages := flag.Int("pages", 1, "Number Of Pages")
	file := flag.String("file", "", "Name Of XLSX File")

	flag.Parse()

	list := used.Fetch(*query, *pages)
	used.WriteXlsFile(*file, list)
}
