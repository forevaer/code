// Sample program to show how to show you how to briefly work with io.
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"code/chapter3/words"
)

//统计文本单词个数
func main() {
	filename := os.Args[1]
	// 读取文件
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}
	// 转化文本
	text := string(contents)
	// 统计
	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
