package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// 初始化加载
func init() {
	// 设置日志输出
	log.SetOutput(os.Stdout)
}

// 主函数入口
func main() {
	// 函数运行
	search.Run("president")
}
