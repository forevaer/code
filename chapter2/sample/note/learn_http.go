package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	fmt.Println(filepath.Abs(filepath.Dir(".")))
}
func main() {
	url := "http://www.baidu.com"
	response, _ := http.Get(url)
	file, _ := os.Open("result.xml")
	defer file.Close()
	reader := bufio.NewReader(response.Body)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Print(s)
	}
}
