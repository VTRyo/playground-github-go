package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func getPublicReposByOrg() {
	client := github.NewClient(nil)

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 1}, // pagenation
		Type:        "public",                       // org type
	}

	repos, _, err := client.Repositories.ListByOrg(context.Background(), "moneyforward", opt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(repos)
}

func main() {
	envErr := godotenv.Load(fmt.Sprintf(".env"))
	if envErr != nil {
		fmt.Printf("環境変数を読み込めませんでした: %v", envErr)
	}

	repoName := ""
	ownerName := "VTRyo"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	issues, _, err := client.Issues.ListByRepo(ctx, ownerName, repoName, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(issues)
	getPublicReposByOrg()
}
