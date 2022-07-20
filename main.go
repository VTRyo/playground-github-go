package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func getPublicOrg() {
	client := github.NewClient(nil)

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "moneyforward", opt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(repos)
}
