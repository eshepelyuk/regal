package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/styrainc/regal/pkg/version"
)

type versionCommandParams struct {
	format string
}

func init() {
	params := versionCommandParams{}

	parseCommand := &cobra.Command{
		Use:   "version [--format=json|pretty]",
		Short: "Print the version of Regal",
		Long:  "Show the version and other build-time metadata for the running Regal binary.",

		PreRunE: func(_ *cobra.Command, args []string) error {
			if params.format == "" {
				params.format = "pretty"
			} else if params.format != "json" && params.format != "pretty" {
				return fmt.Errorf("invalid format: %s", params.format)
			}

			return nil
		},

		Run: func(_ *cobra.Command, args []string) {
			vi := version.New()

			switch params.format {
			case formatJSON:
				e := json.NewEncoder(os.Stdout)
				e.SetIndent("", "  ")
				err := e.Encode(vi)
				if err != nil {
					log.SetOutput(os.Stderr)
					log.Println(err)
					os.Exit(1)
				}
			case formatPretty:
				os.Stdout.WriteString(vi.String())
			default:
				log.SetOutput(os.Stderr)
				log.Printf("invalid format: %s\n", params.format)
				os.Exit(1)
			}
		},
	}
	parseCommand.Flags().StringVar(
		&params.format,
		"format",
		formatPretty,
		fmt.Sprintf("Output format. Valid values are '%s' and '%s'.", formatPretty, formatJSON),
	)
	RootCommand.AddCommand(parseCommand)
}
