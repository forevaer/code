package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, err := regexp.MatchString(".*?@.*?\\.com", "664473032@qq.com")
	if err == nil {
		fmt.Println(match)
	} else {
		fmt.Println(err)
	}
}
