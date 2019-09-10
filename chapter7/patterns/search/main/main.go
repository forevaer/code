package main

import (
	"code/chapter7/patterns/search"
	"log"
)

func main() {
	//options
	results := search.Submit(
		"golang",
		search.OnlyFirst,
		// register google
		search.Google,
		// register bing
		search.Bing,
		// register yahoo
		search.Yahoo,
	)

	for _, result := range results {
		log.Printf("\nmain : Results : Info : %+v\n", result)
	}

	//ops
	//results := search.Submit(
	//	"golang",
	//	search.Google,
	//	search.Bing,
	//	search.Yahoo,
	//)
	//
	//for _, result := range results {
	//	log.Printf("main : Results : Info : %+v\n", result)
	//}
}
