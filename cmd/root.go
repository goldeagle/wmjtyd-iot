package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "wmjtyd-iot",
	Short: "wmjtyd-iot backend cli tools",
	Long:  `wmjtyd-iot backend cli tools for IOT system`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello wmjtyd-iot!")
	},
}

func init() {
	cobra.OnInitialize()
}
