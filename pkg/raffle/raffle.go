package raffle

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/go-github/v45/github"
)

func GetAllIssues(client *github.Client, user, repo string) ([]*github.Issue, error) {
	opt := &github.IssueListByRepoOptions{
		ListOptions: github.ListOptions{PerPage: 50},
	}

	var allIssues []*github.Issue
	for {
		issues, resp, err := client.Issues.ListByRepo(context.Background(), user, repo, opt)
		if err != nil {
			return nil, err
		}
		allIssues = append(allIssues, issues...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allIssues, nil
}

func GetShuffledUsersSet(issues []*github.Issue) []string {
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
