/*
Copyright © 2023 Shazahanul Islam Shohag shohag121@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/shohag121/LetMeKnow/github"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Check if you are logged in",
	Long: `This command checks if you are logged in.
		If you are not logged in, it will display your GitHub Information`,
	Run: func(cmd *cobra.Command, args []string) {
		if auth, err := github.IsAuthenticated(); !auth || err != nil {
			fmt.Println("Please login first, using 'letmeknow auth login -t YOURAUTHTOKEN'")
			return
		}
		me, err := github.WhoAmI()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(me))
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
