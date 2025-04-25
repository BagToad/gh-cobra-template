package whoami

import (
	"fmt"

	"github.com/BagToad/gh-cobra-template/internal"
	"github.com/MakeNowJust/heredoc"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)

type whoamiOptions struct {
	greeting string
}

func NewWhoamiCmd(f *internal.Factory, runF func(*whoamiOptions) error) *cobra.Command {
	// This is the subcommand for the gh-cobra-template extension.
	// You can add your subcommand logic here.
	// For example, you might want to parse flags or set up subcommands.

	opts := &whoamiOptions{}

	subCmd := &cobra.Command{
		Use:   "whoami",
		Short: "Example subcommand for the gh-cobra-template extension",
		Long: heredoc.Docf(`
			%[1]ssubcmd%[1]s is an example subcommand for the gh-cobra-template extension.
		`, "`"),
		Example: heredoc.Doc(`
			# Simple greeting
			$ gh cobra-template subcmd

			# Personalized greeting
			$ gh cobra-template subcmd $(whoami)
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if runF != nil {
				return runF(opts)
			}
			return whoamiRun(opts)
		},
	}

	subCmd.Flags().StringVarP(&opts.greeting, "greeting", "g", "Hello", "Greeting message")

	return subCmd
}

func whoamiRun(opts *whoamiOptions) error {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return err
	}

	response := struct{ Login string }{}
	err = client.Get("user", &response)
	if err != nil {
		return err
	}

	fmt.Printf("running as %s\n", response.Login)
	return nil
}
