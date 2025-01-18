package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	blogTitles, err := getLatestBlogTitles("https://golangcode.com")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Blog Titles: ")
	fmt.Printf(blogTitles)
}

func getLatestBlogTitles(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	fmt.Println(doc)
	if err != nil {
		return "", err
	}

	titles := ""
	// this loops through all the css class selector
	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
		titles += "-" + s.Text() + "\n"
	})

	// to avoid resource leak
	resp.Body.Close()
	return titles, nil
}
