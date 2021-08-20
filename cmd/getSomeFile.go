//+build local

package cmd

import (
	"github.com/discoriver/building-tool-suite/tools"
	"github.com/spf13/cobra"
)

var (
	getSomeFileFlags tools.GetSomeFileCommandFlag

	getSomeFileCmd = &cobra.Command{
		Use:   "getsomefile",
		Short: "Print out the bytes of the provided file.",
		Run: func(cmd *cobra.Command, args []string) {
			getSomeFileFlags.GetSomeFile()
		},
	}
)

func init() {
	rootCmd.AddCommand(getSomeFileCmd)

	getSomeFileCmd.Flags().StringVarP(&getSomeFileFlags.FilePath, "file", "f", "", "Path to file")

	// Required flags.
	getSomeFileCmd.MarkFlagRequired("file")
}
