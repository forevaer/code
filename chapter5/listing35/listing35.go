package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 入口
func main() {
	var b bytes.Buffer
	// 写入
	b.Write([]byte("Hello"))
	// 追加到输出流
	_, _ = fmt.Fprintf(&b, "World!")
	//   对接到输出
	_, _ = io.Copy(os.Stdout, &b)
}
