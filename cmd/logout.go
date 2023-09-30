/*
Copyright © 2023 Shazahanul Islam Shohag shohag121@gmail.com
*/
package cmd

import (
	"fmt"
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
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("error writing config", err)
			return
		}
		fmt.Println("Logged out successfully")
	},
}

func init() {
	authCmd.AddCommand(logoutCmd)
	var force = false
	logoutCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "force logout")
	viper.BindPFlag("force", logoutCmd.PersistentFlags().Lookup("force"))
}
