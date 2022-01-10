package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

const collyLinc = "https://stackoverflow.com/questions/67203641/missing-go-sum-entry-for-module-providing-package-package-name"

func main() {
	res, err := http.Get(collyLinc)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res.Body)

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if res.StatusCode != 200 {
		log.Println("not 200")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}

	// Find the review items
	fmt.Println(doc.Find("#price-value > span > span.js-item-price").Text())
}