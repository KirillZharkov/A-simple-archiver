package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Simple archiver",
}

// the function runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleErr(err)
	}
}
func handleErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
