package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goodmorning",
	Short: "Goodmorning helps you automate you morning routine",
	Long: `A Fast and Flexible morning routine automater built with BubbleTea from Charm.`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Goodmorning")
	},
  }
  
  func Execute() {
	rootCmd.AddCommand(settingsCmd)
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }
  