/*
Copyright © 2023 Shazahanul Islam Shohag shohag121@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/shohag121/LetMeKnow/github"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all unread notifications.",
	Long:  `List all the notifications from GitHub. You can use this command to list all the unread notifications.`,
	Run: func(cmd *cobra.Command, args []string) {
		list, err := github.GetUserNotifications()

		if err != nil {
			fmt.Println(err)
		}

		// TODO: Format and print the list
		fmt.Println(list)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().String("foo", "", "A help for foo")
}
