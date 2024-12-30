package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "notes",
	Short: "Add todo notes to desktop gui",
	Long:  "Add todo notes to desktop gui. Using other libraries that paint the gui you can see your notes directly imported from local file or external repo",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while executing notes", err)
		os.Exit(1)
	}
}
