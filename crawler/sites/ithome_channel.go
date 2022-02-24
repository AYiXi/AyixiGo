package sites

import (
	"ayixigo/crawler/model"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var ch = make(chan *ItHomeItem)

func ItHomeParserChan() {
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
			Title:   title,
			Url:     url,
			Date:    date,
		}

		go GetContentChan(url, &item, i)

		fmt.Println(item)
	})
}

func GetContentChan(url string, item *ItHomeItem, i int) {
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
	ch <- item
}

func Insert() {
	for {
		item, ok := <-ch
		if ok {
			// if err := model.DB.Table(tableName).Create(item).Error; err != nil {
			// 	log.Printf("db.Create, err : %v", err)
			// }
			model.DB.Table(tableName).Create(item)
		}
	}
}

func ItHomeSpiderChan() {
	defer close(ch)
	go ItHomeParserChan()
	go Insert()

	time.Sleep(time.Second * 2)
	// select {
	// 	case <- time.After(time.Second * 1):
	// 		if err := model.DB.Table(tableName).Create(item).Error; err != nil {
	// 			log.Printf("db.Create err : %v", err)
	// 		}
	// 	case <- ch:

	// }
}
