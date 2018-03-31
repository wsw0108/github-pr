package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

var (
	user  string
	repo  string
	head  string
	base  string
	title string
)

func init() {
	flag.StringVar(&user, "user", "", "repo owner")
	flag.StringVar(&repo, "repo", "", "repo name")
	flag.StringVar(&head, "head", "", "head to create pull request")
	flag.StringVar(&base, "base", "", "base to create pull request")
	flag.StringVar(&title, "title", "", "title for pull request")
	flag.Parse()
}

func main() {
	if user == "" || repo == "" || head == "" || base == "" {
		flag.Usage()
		os.Exit(1)
	}

	accessToken, ok := os.LookupEnv("GITHUB_ACCESS_TOKEN")
	if !ok {
		fmt.Println("Please set 'GITHUB_ACCESS_TOKEN' first.")
		os.Exit(-1)
	}

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	pull := &github.NewPullRequest{
		Title: &title,
		Head:  &head,
		Base:  &base,
	}

	pr, _, err := client.PullRequests.Create(ctx, user, repo, pull)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("PR created at %s", *pr.URL)
	os.Exit(0)
}
