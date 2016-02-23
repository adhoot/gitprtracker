package main

import (
	"net/http"
	"log"
	"gitprtracker/ui"
	"gitprtracker/gitapi"
	"fmt"
)

func main() {
	mainHandler := ui.MainPage

	prs, err := gitapi.SearchPRs()
	if (err != nil) {
		log.Fatal("Error getting search results")
	}

	fmt.Printf("Found %d issues", prs.TotalCount)

	for _, pr := range prs.Items {
		fmt.Printf("%v \n", pr)
	}

	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
	println("Git pull request tracking started ...")
}
