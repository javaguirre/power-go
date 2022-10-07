package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

func main() {
	// TODO Get repository list
	// TODO Goroutine to search for the integrations functions
	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	integrationMethods := []string{
		"checkIntegrationLimitsForConfigSave",
		"getInstalledIntegrations",
		"checkIfIntegrationsMeetFreemiumLimits",
	}

	client := github.NewClient(tc)
	searchOptions := github.SearchOptions{}

	// List repositories
	resultRepos, response, err := client.Search.Repositories(
		ctx,
		"mattermost-plugin in:name",
		&searchOptions,
	)
	if err != nil || response.StatusCode != 200 {
		fmt.Println("Errror! %q", err)
		fmt.Println(response)
	}

	totalReposCount := resultRepos.GetTotal()
	repoNames := []string{}
	totalResults := 0
	reposErrored := 0

	for _, repo := range resultRepos.Repositories {
		if *repo.Name == "" {
			continue
		}
		repoNames = append(repoNames, *repo.Name)
	}

	resultSearch := []string{}

	for _, name := range repoNames {
		fmt.Printf("searching on %s...\n", name)

		query := fmt.Sprintf("%s+repo:mattermost/%s", integrationMethods[0], name)
		result, _, err := client.Search.Code(ctx, query, &searchOptions)
		if err != nil {
			// TODO Get the ones with error HTTP 422
			reposErrored++
			fmt.Println("Errror! %q", err)
		}

		if result.GetTotal() > 0 {
			// TODO Get which search
			totalResults += result.GetTotal()
			resultSearch = append(resultSearch, name)
		}

		fmt.Printf("%d Results\n", result.GetTotal())
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("Total results :%d\n", totalResults)
	fmt.Printf("Total repos :%d\n", totalReposCount)
	fmt.Printf("Total repos errored:%d\n", reposErrored)
	for _, item := range resultSearch {
		fmt.Println(item)
	}
}
