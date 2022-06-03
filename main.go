package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/go-github/v45/github"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./main <user> <repo>")
		os.Exit(1)
	}

	user := os.Args[1]
	repo := os.Args[2]

	client := github.NewClient(nil)

	allIssues, err := getAllIssues(client, user, repo)
	if err != nil {
		panic(err)
	}
	if len(allIssues) == 0 {
		fmt.Println("No issues found")
		os.Exit(0)
	}

	shuffledUsers := getShuffledUsersSet(allIssues)

	fmt.Println(len(allIssues), shuffledUsers[0])
}

func getAllIssues(client *github.Client, user, repo string) ([]*github.Issue, error) {
	opt := &github.IssueListByRepoOptions{
		ListOptions: github.ListOptions{PerPage: 50},
	}

	var allIssues []*github.Issue
	for {
		issues, resp, err := client.Issues.ListByRepo(context.Background(), user, repo, opt)
		if err != nil {
			panic(err)
		}
		allIssues = append(allIssues, issues...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allIssues, nil
}

func getShuffledUsersSet(issues []*github.Issue) []string {
	usersMap := make(map[string]struct{})
	for _, issue := range issues {
		if issue.GetUser().Login != nil {
			usersMap[*issue.GetUser().Login] = struct{}{}
		}
	}

	users := []string{}
	for user := range usersMap {
		users = append(users, user)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(users), func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})

	return users
}
