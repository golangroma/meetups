package cli

import (
	"fmt"
	"strings"

	"github.com/golangroma/meetup-20220614/pkg/raffle"
	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "raffle",
		Short:         "raffle issue owners!",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	rootCmd.AddCommand(NewRunCommand())
	rootCmd.AddCommand(NewListCommand())

	return rootCmd
}

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:           "list",
		Short:         "list partecipants",
		Args:          UserRepoArg(),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			userRepo := strings.Split(args[0], "/")
			user, repo := userRepo[0], userRepo[1]

			client := github.NewClient(nil)
			allIssues, err := raffle.GetAllIssues(client, user, repo)
			if err != nil {
				return err
			}

			shuffledUsers := raffle.GetShuffledUsersSet(allIssues)
			for _, user := range shuffledUsers {
				fmt.Println(user)
			}

			return nil
		},
	}
}

func NewRunCommand() *cobra.Command {
	return &cobra.Command{
		Use:           "run",
		Short:         "run raffle",
		Args:          UserRepoArg(),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			userRepo := strings.Split(args[0], "/")
			user, repo := userRepo[0], userRepo[1]

			client := github.NewClient(nil)
			allIssues, err := raffle.GetAllIssues(client, user, repo)
			if err != nil {
				return err
			}

			shuffledUsers := raffle.GetShuffledUsersSet(allIssues)
			fmt.Println(shuffledUsers[0])

			return nil
		},
	}
}

func UserRepoArg() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("accepts only one argument OWNER/REPO")
		}

		userRepo := strings.Split(args[0], "/")
		if len(userRepo) != 2 {
			return fmt.Errorf("accepts only one argument OWNER/REPO")
		}

		return nil
	}
}
