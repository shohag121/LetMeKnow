/*
Copyright © 2023 Shazahanul Islam Shohag shohag121@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/shohag121/LetMeKnow/github"
	ghuser "github.com/shohag121/LetMeKnow/user"
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
		resp, err := github.WhoAmI()

		if err != nil {
			fmt.Println(err)
		}

		var user ghuser.User
		err = json.Unmarshal(resp, &user)
		if err != nil {
			fmt.Println(err)
		}
		ghuser.Format(user)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
