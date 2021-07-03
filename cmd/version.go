package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd *cobra.Command

func init() {
	versionCmd = &cobra.Command{
		Use: "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("goblockbastards 1.0")
		},
	}
	rootCmd.AddCommand(versionCmd)
}
