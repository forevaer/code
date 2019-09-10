package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	//
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	// append 2 buffer
	fmt.Fprintf(&b, "World!")

	b.WriteTo(os.Stdout)
}
