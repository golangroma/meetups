package main

import (
	"fmt"

	"github.com/golangroma/meetup-20220614/pkg/cli"
)

func main() {
	rootCmd := cli.NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
