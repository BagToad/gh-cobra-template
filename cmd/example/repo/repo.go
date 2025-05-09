package repo

import (
	"fmt"

	"github.com/BagToad/gh-cobra-template/internal"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// repoOptions is intentionally empty because this example is demonstrating
// how the persistent `-R, --repo` flag could be used in a commandset.
type repoOptions struct{}

func NewRepoCmd(f *internal.Factory, runF func(*repoOptions) error) *cobra.Command {
	// This is the details command for the gh-cobra-template extension.
	// You can add your command logic here.
	// For example, you might want to parse flags or set up subcommands.

	opts := &repoOptions{}

	detailsCmd := &cobra.Command{
		Use:   "repo",
		Short: "Details command for the gh-cobra-template extension",
		Long: heredoc.Doc(`

		`),
		Example: heredoc.Doc(`

		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if runF != nil {
				return runF(opts)
			}
			return detailsRun(f, opts)
		},
	}

	return detailsCmd
}

func detailsRun(f *internal.Factory, _ *repoOptions) error {
	repo, err := f.RepositoryResolver()
	if err != nil {
		return err
	}

	fmt.Printf("The repository detected is: %s/%s\n", repo.Owner, repo.Name)
	return nil
}
