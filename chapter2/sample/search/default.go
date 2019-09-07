package search

// 默认匹配器
type defaultMatcher struct {
}

// 初始化时自定注册
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// 默认返回
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
