package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var settingsCmd = &cobra.Command{
	Use: "settings",
	Short: "Manage your routine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("settings")
	},
}