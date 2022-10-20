package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewSettingsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "settings",
		Short: "Manage your routine",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("settings")
			return nil
		},
	}
}
