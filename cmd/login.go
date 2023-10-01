/*
Copyright © 2023 Shazahanul Islam Shohag shohag121@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/ggwhite/go-masker"
	"github.com/shohag121/LetMeKnow/cron"
	"github.com/shohag121/LetMeKnow/github"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:     "login",
	Short:   "Login by providing your Github access token",
	Long:    `Login by providing your Github access token. eg: letmeknow login -t <token>`,
	Example: `letmeknow login -t <token>`,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		fmt.Println("Checking token...")
		fmt.Println("Token Provided:", masker.String(masker.MAddress, token))
		viper.Set("token", token)
		if auth, err := github.IsAuthenticated(); !auth || err != nil {
			fmt.Println("You are not authenticated. Please provide a valid token.")
			return
		}
		fmt.Println("Login Successful!")
		viper.Set("authenticated", true)

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("We have problem saving your token.", err)
			return
		}
		cron.AddCronJob()
	},
}

func init() {
	authCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringP("token", "t", "", "Github access token")
	loginCmd.MarkPersistentFlagRequired("token")
	viper.BindPFlag("token", loginCmd.PersistentFlags().Lookup("token"))
}
