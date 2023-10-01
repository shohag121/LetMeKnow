/*
Copyright © 2023 Shazahanul Islam Shohag shohag121@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/shohag121/LetMeKnow/github"
	"github.com/shohag121/LetMeKnow/notification"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all unread notifications.",
	Long: `List all the notifications from GitHub.
			You can use this command to list all the unread notifications.`,
	Run: func(cmd *cobra.Command, args []string) {
		list, err := github.GetUserNotifications()

		if err != nil {
			fmt.Println(err)
		}

		if viper.GetBool("display") {
			notification.Display(list)
		} else {
			notification.Process(list)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().BoolP("display", "d", true, "Display as list")
	viper.BindPFlag("display", listCmd.PersistentFlags().Lookup("display"))
}
