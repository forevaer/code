package main

import (
	_ "code/chapter2/sample/matchers"
	"code/chapter2/sample/search"
	"log"
	"os"
)

// 初始化加载
func init() {
	// 设置日志输出
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Llongfile)
}

// 主函数入口
func main() {
	// 函数运行
	search.Run("president")
}
