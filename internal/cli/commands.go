package cli

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "words",
		Short: "Word form replacement tool",
	}
}

func NewReplaceCommand(cli *CLI) *cobra.Command {
	return &cobra.Command{}
}
