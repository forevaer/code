package main

import (
	"code/chapter2/sample/search"
	"log"
	"os"
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
