package cmd

import (
	"goodmorning/cmd/settings"

	"github.com/spf13/cobra"
)

func NewSettingsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "settings",
		Short: "Manage your routine",
		RunE: func(cmd *cobra.Command, args []string) error {
			return settings.Program()
		},
	}
}
