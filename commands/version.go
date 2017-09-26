package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ReleaseVersion string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of dds",
	Long:  `All software has versions. This is DDS's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("DDS Version: %s\n", ReleaseVersion)
	},
}
