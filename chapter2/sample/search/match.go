package search

import (
	"log"
)

// 查询结果
type Result struct {
	Field   string
	Content string
}

// 匹配器
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 调用并返回
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 把结果传入通道
	for _, result := range searchResults {
		results <- result
	}
}

// 展示
func Display(results chan *Result) {
	// 把查询的结果进行展示
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
