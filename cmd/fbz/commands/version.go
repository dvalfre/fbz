package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string
	Build   string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Report the fbz version in use",
	Long: `Report the fbz version in use

The version and build time are set during our release build, and this can
help you to guarantee that you're using the most recent version.`,
	Run: func(cmd *cobra.Command, args []string) {
		showVersion()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func showVersion() {
	fmt.Println("fbz", Version, "(Build "+Build+")")
}
