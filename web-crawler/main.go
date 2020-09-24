package main

import (
	"context"
	"fmt"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func main() {
	ctx := context.Background()
	result_data, _ := googlesearch.Search(ctx, "8901058101614")
	for _, titles := range result_data {
		fmt.Println(titles.Title)
	}
}
