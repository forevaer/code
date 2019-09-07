package search

import (
	"encoding/json"
	"log"
	"os"
)

const dataFile = "./data/data.json"

// 结构定义
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// 读取文件
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		panic(err)
		return nil, err
	}
	// 关闭文件
	defer file.Close()
	// 结构读取
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	// 返回
	return feeds, err
}

func Show(file *os.File) {
	stat, err := file.Stat()
	if err == nil {
		length := stat.Size()
		bs := make([]byte, length)
		_, _ = file.Read(bs)
		log.Println(string(bs))
	} else {
		panic(err)
	}
}
