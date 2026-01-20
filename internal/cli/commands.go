package cli

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "words",
		Short: "Word form replacement tool",
	}
}

func NewReplaceCommand(cli *CLI) *cobra.Command {
	var inputFile, dataFile, oldWord, newWord string

	cmd := &cobra.Command{
		Use:   "replace",
		Short: "Replace word forms in text",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.RunWithFiles(inputFile, dataFile, oldWord, newWord)
		},
	}

	cmd.Flags().StringVar(&inputFile, "input", "", "Input file to process")
	cmd.Flags().StringVar(&dataFile, "data", "", "JSONL data file")
	cmd.Flags().StringVar(&oldWord, "old", "", "Word to replace")
	cmd.Flags().StringVar(&newWord, "new", "", "New word")

	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("old")
	cmd.MarkFlagRequired("new")

	return cmd
}
