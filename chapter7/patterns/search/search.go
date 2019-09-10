package search

import (
	"fmt"
	"log"
)

type Result struct {
	Engine      string
	Title       string
	Description string
	Link        string
}

type Searcher interface {
	Search(searchTerm string, searchResults chan<- []Result)
}

type searchSession struct {
	searchers  map[string]Searcher
	first      bool
	resultChan chan []Result
}

func Google(s *searchSession) {
	log.Println("search : Submit : Info : Adding Google")
	s.searchers["google"] = google{}
}

func Bing(s *searchSession) {
	log.Println("search : Submit : Info : Adding Bing")
	s.searchers["bing"] = bing{}
}

func Yahoo(s *searchSession) {
	log.Println("search : Submit : Info : Adding Yahoo")
	s.searchers["yahoo"] = yahoo{}
}

// set first
func OnlyFirst(s *searchSession) {
	s.first = true
}

func Submit(query string, options ...func(*searchSession)) []Result {
	var session searchSession
	session.searchers = make(map[string]Searcher)
	session.resultChan = make(chan []Result)
	// 前置
	// execute funcs
	// init and register
	for _, opt := range options {
		opt(&session)
	}

	// search
	for _, s := range session.searchers {
		// 然后写到resultChan
		go s.Search(query, session.resultChan)
	}

	var results []Result
	// 全部的搜索器
	for search := 0; search < len(session.searchers); search++ {
		// 每个搜索器一个结果
		if session.first && search > 0 {
			go func() {
				// 查询结果到r
				r := <-session.resultChan
				log.Printf("search : Submit : Info : Results Discarded : Results[%d]\n", len(r))
			}()
			continue
		}
		fmt.Println("========================")
		fmt.Println(session.first, search)
		fmt.Println("========================")
		// Wait to recieve results.
		log.Println("search : Submit : Info : Waiting For Results...")
		// wait for result
		result := <-session.resultChan

		// Save the results to the final slice.
		log.Printf("search : Submit : Info : Results Used : Results[%d]\n", len(result))
		// 添加到队列
		results = append(results, result...)
	}
	// complete
	log.Printf("search : Submit : Completed : Found [%d] Results\n", len(results))
	return results
}
