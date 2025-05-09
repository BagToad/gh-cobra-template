// example package is used to demonstrate how the GitHub CLI maintainers structure
// Cobra commands and application logic for testing purposes.
//
// This is meant to adapted or deleted as you see fit.
package example

import (
	"github.com/BagToad/gh-cobra-template/cmd/example/repo"
	"github.com/BagToad/gh-cobra-template/cmd/example/whoami"
	"github.com/BagToad/gh-cobra-template/internal"
	"github.com/spf13/cobra"
)

func NewExampleCmd(f *internal.Factory) *cobra.Command {
	// This is the example command for the gh-cobra-template extension.
	// You can add your command logic here.
	// For example, you might want to parse flags or set up subcommands.

	exampleCmd := &cobra.Command{
		Use:   "example",
		Short: "Example command for the gh-cobra-template extension",
		Long:  "This is an example command for the gh-cobra-template extension.",
	}

	exampleCmd.AddCommand(repo.NewRepoCmd(f, nil))
	exampleCmd.AddCommand(whoami.NewWhoamiCmd(f, nil))

	return exampleCmd
}
