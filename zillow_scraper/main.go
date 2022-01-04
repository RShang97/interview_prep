package main

import (
	"fmt"

	"github.com/RShang97/interview_prep/zillow_scraper/zillow_structs"
)

func main() {
	fmt.Println("hello world")

	searchResultsRequest := zillow_structs.GetSearchResultsRequest{
		ZwsID:        "X1-ZWz16e4zj689or_8s1xw",
		Address:      "420 Pontius Ave N",
		CityStateZip: "Seattle, WA",
	}
	response, err := zillow_structs.GetSearchResults(&searchResultsRequest)
	fmt.Println(response, err)
}
