//+build jump

package cmd

import (
	"github.com/discoriver/building-tool-suite/tools"
	"github.com/spf13/cobra"
)

var (
	parseSomeDBQueryFlags tools.ParseSomeDBQueryCommandFlag

	parseSomeDBQueryCmd = &cobra.Command{
		Use:   "parsesomequery",
		Short: "Print out the bytes of the provided file.",
		Run: func(cmd *cobra.Command, args []string) {
			parseSomeDBQueryFlags.ParseSomeDBQuery()
		},
	}
)

func init() {
	rootCmd.AddCommand(parseSomeDBQueryCmd)

	parseSomeDBQueryCmd.Flags().StringVarP(&parseSomeDBQueryFlags.ID, "ID", "i", "", "ID filter for query")

	// Required flags.
	parseSomeDBQueryCmd.MarkFlagRequired("ID")
}
