package search

import (
	"log"
	"sync"
)

// 匹配器容器
var matchers = make(map[string]Matcher)

// 直接匹配
func Run(searchTerm string) {
	// 读取文件
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// 结果合集
	results := make(chan *Result)
	// 调度
	var waitGroup sync.WaitGroup
	// 一次性直接添加任务
	waitGroup.Add(len(feeds))
	// 异步遍历
	for _, feed := range feeds {
		// 查找匹配器
		matcher, exists := matchers[feed.Type]
		// 默认匹配器
		if !exists {
			matcher = matchers["default"]
		}
		// 匹配动作
		go func(matcher Matcher, feed *Feed) {
			// 匹配
			Match(matcher, feed, searchTerm, results)
			// 通知
			waitGroup.Done()
		}(matcher, feed)
	}
	// 异步
	go func() {
		// 等待全部完成
		waitGroup.Wait()

		// 完成以后关闭通道
		close(results)
	}()

	// 显示结果
	Display(results)
}

// 注册匹配器
func Register(feedType string, matcher Matcher) {
	// 不允许重名注册
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
