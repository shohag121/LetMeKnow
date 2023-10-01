/*
Copyright © 2023 Shazahanul Islam Shohag shohag121@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/shohag121/LetMeKnow/cron"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:     "logout",
	Short:   "Logout from the GitHub",
	Long:    ``,
	Example: `letmeknow auth logout -f`,
	Run: func(cmd *cobra.Command, args []string) {
		force := viper.GetBool("force")
		if !force {
			fmt.Println("Are you sure you want to logout? use `-f` to force logout.")
			return
		}

		viper.Set("authenticated", false)
		viper.Set("token", "")
		viper.Set("force", false)
		viper.Set("last_notifications", "")
		viper.Set("last_result_notifications", "")
		viper.Set("last_user", "")
		viper.Set("last_result_user", "")

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("error writing config", err)
			return
		}
		fmt.Println("Logged out successfully")
		cron.RemoveCronJob()
	},
}

func init() {
	authCmd.AddCommand(logoutCmd)
	var force = false
	logoutCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "force logout")
	viper.BindPFlag("force", logoutCmd.PersistentFlags().Lookup("force"))
}
