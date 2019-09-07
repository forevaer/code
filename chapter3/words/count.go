package words

import "strings"

// 统计单词字母
func CountWords(text string) int {
	return len(strings.Fields(text))
}
