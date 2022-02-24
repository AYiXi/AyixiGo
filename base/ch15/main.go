package main

import (
	"fmt"
	"ayixigo/base/ch15/series"

	"github.com/PuerkitoBio/goquery"
	"rsc.io/quote"
)

func main() {
	fmt.Println("1")
	fmt.Println(series.GetFSeries(10))
	quote.Hello()
	doc := new(goquery.Document)
	fmt.Println(doc)
	
}
