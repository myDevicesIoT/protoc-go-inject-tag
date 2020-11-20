package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	var inputFiles string
	var xxxTags string
	flag.StringVar(&inputFiles, "input", "", "path to input file")
	flag.StringVar(&xxxTags, "XXX_skip", "", "skip tags to inject on XXX fields")
	flag.BoolVar(&verbose, "verbose", false, "verbose logging")

	flag.Parse()

	var xxxSkipSlice []string
	if len(xxxTags) > 0 {
		xxxSkipSlice = strings.Split(xxxTags, ",")
	}

	if len(inputFiles) == 0 {
		log.Fatal("input file is mandatory")
	}

	for _, inputFile := range strings.Split(inputFiles, " ") {
		areas, err := parseFile(inputFile, xxxSkipSlice)
		if err != nil {
			log.Fatal(err)
		}
		if err = writeFile(inputFile, areas); err != nil {
			log.Fatal(err)
		}
	}
}
