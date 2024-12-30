package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var external bool
var addNoteCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"post"},
	Short:   "Add note to desktop",
	Long:    "Add local note to desktop",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var res string
		var err error
		if external {
			res, err = addRemoteFile(args[0])
		} else {
			res, err = addFile(args[0])
		}
		if err != nil {
			fmt.Printf("Error adding note: ", err)
			return err
		}
		fmt.Printf("Success adding note to desktop", res)
		return nil
	},
}

func init() {
	addNoteCmd.Flags().BoolVarP(&external, "external", "e", false, "Fetch note from remote url")
	rootCmd.AddCommand(addNoteCmd)
}
