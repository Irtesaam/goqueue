package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long: `Display version information for GoQueue including:
- Version number
- Git commit hash
- Build date`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("GoQueue %s\n", Version)
		if GitCommit != "unknown" {
			fmt.Printf("Git commit: %s\n", GitCommit)
		}
		if BuildDate != "unknown" {
			fmt.Printf("Built: %s\n", BuildDate)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
