package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of wmjtyd-iot and Go environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wmjtyd-iot Version:", "1.0.0")
		fmt.Println("Go Version:", runtime.Version())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
