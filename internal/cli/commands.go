package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bryack/words/adapters/sqlite"
	"github.com/bryack/words/internal/replacer"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "words",
		Short: "Word form replacement tool",
	}
}

func NewReplaceCommand() *cobra.Command {
	var inputFile, dataFile, oldWord, newWord string

	cmd := &cobra.Command{
		Use:   "replace",
		Short: "Replace word forms in text",
		RunE: func(cmd *cobra.Command, args []string) error {
			configDir, err := os.UserConfigDir()
			if err != nil {
				return err
			}
			dbDir := filepath.Join(configDir, "words")
			if err := os.MkdirAll(dbDir, 0755); err != nil {
				return fmt.Errorf("failed to create config directory %s: %w", dbDir, err)
			}
			dbPath := filepath.Join(dbDir, "words.db")
			var loader sqlite.DataLoader
			if dataFile != "" {
				loader = sqlite.LoadFromJSONLFile(dataFile)
			}

			provider, err := sqlite.NewSQLiteFormsProvider(dbPath, loader)
			if err != nil {
				return fmt.Errorf("failed to create provider with database %s: %w", dbPath, err)
			}
			defer provider.Close()

			replacer := replacer.NewReplacer(provider)
			cli := NewCLI(os.Stdin, os.Stdout, replacer)

			return cli.RunWithFiles(inputFile, oldWord, newWord)
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
