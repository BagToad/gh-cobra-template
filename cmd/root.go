package cmd

import (
	"github.com/BagToad/gh-cobra-template/cmd/example"
	"github.com/BagToad/gh-cobra-template/internal"
	"github.com/MakeNowJust/heredoc"
	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/spf13/cobra"
)

func NewRootCmd(f *internal.Factory) *cobra.Command {
	// This is the main entry point for the gh-cobra-template extension.
	// You can add your base command logic here.
	// For example, you might want to parse flags or set up subcommands.

	rootCmd := &cobra.Command{
		Use:   "gh-cobra-template",
		Short: "gh-cobra-template is a GitHub CLI extension for testing kwaf",
		Long: heredoc.Doc(`
			gh-cobra-template is a GitHub CLI extension template for demonstrating
			an opinionated implementation using Cobra and following cli/cli conventions.

			This is the long form description of the gh-cobra-template extension.
		`),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// This is the value the user provided with the -R / --repo flag
			repoOverride, _ := cmd.Flags().GetString("repo")
			f.RepositoryResolver = func() (repository.Repository, error) {
				// If the user provided a repository with -R / --repo, use that
				if repoOverride != "" {
					return repository.Parse(repoOverride)
				}
				// Otherwise, infer the repository from the CWD
				return repository.Current()
			}
		},
	}

	rootCmd.PersistentFlags().StringP("repo", "R", "", "Specify the repository to use")

	rootCmd.AddCommand(example.NewExampleCmd(f))

	return rootCmd
}
