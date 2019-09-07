package main

import (
	"fmt"
	"strings"
)

func main() {
	source := "   first second third"
	spaceSplit := strings.Fields(source)
	fmt.Println(spaceSplit)

	customerSpliter := func(rs ...rune) func(r rune) bool {
		return func(r rune) bool {
			for _, rr := range rs {
				if rr == r {
					return true
				}
			}
			return false
		}
	}
	customerResult := strings.FieldsFunc(source, customerSpliter('t', 'n', 'r', ' '))
	fmt.Println(customerResult)
}
