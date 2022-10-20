package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() (*cobra.Command) {
	cmd :=&cobra.Command{
		Use:   "goodmorning",
		Short: "Goodmorning helps you automate you morning routine",
		Long:  `A Fast and Flexible morning routine automater built with BubbleTea from Charm.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(),"Goodmorning")
		},
	}
	
	cmd.AddCommand(NewSettingsCmd())
	return cmd
}

func Execute() {
	rootCmd := NewRootCmd()
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
