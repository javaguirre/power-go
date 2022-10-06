package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v47/github"
)

func main() {
	client := github.NewClient(nil)
	searchOptions := github.SearchOptions{}
	result, _, err := client.Search.Code(
		context.Background(),
		"integrations+repo:mattermost/mattermost-plugin-jira",
		&searchOptions,
	)
	if err != nil {
		fmt.Println("Errror! %q", err)
	}

	fmt.Println("GitHub search on repos")
	fmt.Println(result.GetTotal())

	for _, item := range result.CodeResults {
		fmt.Println(fmt.Sprintf("%s: %s", *item.Name, *item.Path))
	}
}
