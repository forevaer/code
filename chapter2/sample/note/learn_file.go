package main

import (
	"log"
	"os"
)

// 初始化
func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	// 文件路径
	filePath := "data/data.json"
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
		return
	}
	// 关闭
	defer file.Close()
	stat, err := file.Stat()
	length := stat.Size()
	log.Println("size : ", length)
	// 容器
	bs := make([]byte, length)
	// 接受
	_, _ = file.Read(bs)
	println(string(bs))
}
