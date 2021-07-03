package cmd

import (
	"github.com/spf13/cobra"

	"fmt"
	"os"
)


var rootCmd = &cobra.Command{
  Use: "goblockbastards",
  SilenceUsage: true,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize()
}
