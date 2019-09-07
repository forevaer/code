package note

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	for index, data := range arr {
		fmt.Printf("%d : %d \n", index, data)
	}
}
