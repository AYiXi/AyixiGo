package sites

import (
	"ayixigo/crawler/model"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var (
	url       = "https://www.ithome.com/list/"
	tableName = "crawler_ithome"
)

type ItHomeItem struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

func ItHomeParser() []*ItHomeItem {
	var items []*ItHomeItem
	var wg sync.WaitGroup

	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#list > div.fl > ul.datel > li").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a.t").Text()
		url, _ := s.Find("a.t").Attr("href")
		date := s.Find("i").Text()

		if strings.Contains(url, "lapin") {
			fmt.Println("广告 url ->", url)
			return
		}

		fmt.Println("i :", i)

		item := ItHomeItem{
			Title: title,
			Url:   url,
			Date:  date,
		}
		wg.Add(1)
		go GetContent(url, &item, &wg, i)

		items = append(items, &item)

		fmt.Println(item)
	})

	wg.Wait()
	return items
}

func GetContent(url string, item *ItHomeItem, wg *sync.WaitGroup, i int) {
	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
	}

	var content string
	doc.Find("#paragraph > p").Each(func(i int, s *goquery.Selection) {
		content += s.Text()
	})

	if len(content) > 0 {
		fmt.Println(i, content[:20])
	} else {
		fmt.Println("content is empty:", url)
	}

	item.Content = content
	wg.Done()
}

func ItHomeSpider() {
	docs := ItHomeParser()

	for i, doc := range docs {
		if err := model.DB.Table(tableName).Create(doc).Error; err != nil {
			log.Printf("db.Create index: %d, err : %v", i, err)
		}
	}
}
