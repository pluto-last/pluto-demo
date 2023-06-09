package main

import (
	"github.com/PuerkitoBio/goquery"
	charseter "golang.org/x/net/html/charset"
	"log"
	"net/http"
)

func main() {

	// 抓取深圳学校网上的学校
	resp, err := http.Get("https://www.szxuexiao.com/HighSchool/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 转换编码
	reader, err := charseter.NewReader(resp.Body, "utf-8")
	if err != nil {
		log.Println("编码成utf-8失败")
		return
	}

	// 使用 goquery 解析网页内容
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println("编码成utf-8失败")
		return
	}

	// 通过元素选择器解析 document
	doc.Find(".xuexiaopiclist li").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		log.Println("拿到的学校名称是:", content)
	})
}
