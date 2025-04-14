package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/spf13/cobra"
)

type rootOptions struct {
	repositoryResolver func() (repository.Repository, error)
}

func NewRootCmd() *cobra.Command {
	// This is the main entry point for the gh-cobra-template extension.
	// You can add your base command logic here.
	// For example, you might want to parse flags or set up subcommands.

	opts := &rootOptions{}

	rootCmd := &cobra.Command{
		Use:   "gh-cobra-template",
		Short: "gh-cobra-template is a GitHub CLI extension for testing kwaf",
		Long: heredoc.Doc(`
			gh-cobra-template is a GitHub CLI extension template for demonstrating
			an opinionated implementation using Cobra and following cli/cli conventions.

			This is the long form description of the gh-cobra-template extension.
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			repo, err := opts.repositoryResolver()
			if err != nil {
				return err
			}

			fmt.Printf("The repository detected is: %s/%s\n", repo.Owner, repo.Name)
			return nil

			// fmt.Println("hi world, this is the gh-cobra-template extension!")
			// client, err := api.DefaultRESTClient()
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }
			// response := struct{ Login string }{}
			// err = client.Get("user", &response)
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }
			// fmt.Printf("running as %s\n", response.Login)
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// This is the value the user provided with the -R / --repo flag
			repoOverride, _ := cmd.Flags().GetString("repo")
			opts.repositoryResolver = func() (repository.Repository, error) {
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

	return rootCmd
}
